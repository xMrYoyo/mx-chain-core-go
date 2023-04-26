package receiver

import (
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/closing"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/connection"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsReceiver holds the arguments that are needed for a receiver
type ArgsReceiver struct {
	Uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	Log                      core.Logger
	RetryDurationInSec       int
	BlockingAckOnError       bool
}

type receiver struct {
	payloadHandler           webSockets.PayloadHandler
	uint64ByteSliceConverter connection.Uint64ByteSliceConverter
	log                      core.Logger
	safeCloser               core.SafeCloser
	retryDuration            time.Duration
	blockingAckOnError       bool
}

// NewReceiver will create a new instance of receiver
func NewReceiver(args ArgsReceiver) (*receiver, error) {
	return &receiver{
		log:                      args.Log,
		uint64ByteSliceConverter: args.Uint64ByteSliceConverter,
		retryDuration:            time.Duration(args.RetryDurationInSec) * time.Second,
		blockingAckOnError:       args.BlockingAckOnError,
		safeCloser:               closing.NewSafeChanCloser(),
	}, nil
}

// SetPayloadHandler will set the payload handler
func (r *receiver) SetPayloadHandler(handler webSockets.PayloadHandler) {
	r.payloadHandler = handler
}

// Listen will listen for messages from the provided connection
func (r *receiver) Listen(connection connection.WSConClient) (closed bool) {
	timer := time.NewTimer(r.retryDuration)
	defer timer.Stop()

	for {
		_, message, err := connection.ReadMessage()
		if err == nil {
			r.verifyPayloadAndSendAckIfNeeded(connection, message)
			continue
		}

		_, isConnectionClosed := err.(*websocket.CloseError)
		if !isConnectionClosed {
			if strings.Contains(err.Error(), data.ClosedConnectionMessage) {
				r.log.Info("connection closed by server")
				return true
			}
			if strings.Contains(err.Error(), data.ErrConnectionNotOpened.Error()) {
				return
			}

			r.log.Warn("r.Listen()-> connection problem", "error", err.Error())
		}

		select {
		case <-r.safeCloser.ChanClose():
			return
		case <-timer.C:
		}
	}
}

func (r *receiver) verifyPayloadAndSendAckIfNeeded(connection connection.WSConClient, payload []byte) {
	if len(payload) == 0 {
		r.log.Error("empty payload")
		return
	}

	payloadData, err := r.payloadHandler.HandlePayload(payload)
	r.log.LogIfError(err)
	if err != nil && r.blockingAckOnError {
		return
	}

	r.sendAckIfNeeded(connection, payloadData)
}

func (r *receiver) sendAckIfNeeded(connection connection.WSConClient, payloadData *data.PayloadData) {
	if !payloadData.WithAcknowledge {
		return
	}

	r.waitForAckSignal(connection, payloadData.Counter)
}

func (r *receiver) waitForAckSignal(connection connection.WSConClient, counter uint64) {
	timer := time.NewTimer(r.retryDuration)
	defer timer.Stop()

	counterBytes := r.uint64ByteSliceConverter.ToByteSlice(counter)
	for {
		timer.Reset(r.retryDuration)

		err := connection.WriteMessage(websocket.BinaryMessage, counterBytes)
		if err == nil {
			return
		}

		if !strings.Contains(err.Error(), data.ErrConnectionNotOpened.Error()) {
			r.log.Error("could not write acknowledge message", "error", err.Error(), "retrying in", r.retryDuration)
		}

		select {
		case <-timer.C:
		case <-r.safeCloser.ChanClose():
			return
		}
	}
}

// Close will close the underlying ws connection
func (r *receiver) Close() error {
	defer r.safeCloser.Close()

	err := r.payloadHandler.Close()
	if err != nil {
		r.log.Error("cannot close the operations handler", "error", err)
	}

	return nil
}
