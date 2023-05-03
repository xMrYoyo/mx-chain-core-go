package webSockets

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/atomic"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/marshal"
	outportSenderData "github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// ArgsWebSocketsDriver holds the arguments needed for creating a new webSocketsDriver
type ArgsWebSocketsDriver struct {
	Marshaller      marshal.Marshalizer
	WebsocketSender WebSocketSenderHandler
	Log             core.Logger
}

type webSocketsDriver struct {
	marshalizer      marshal.Marshalizer
	log              core.Logger
	payloadConverter PayloadConverter
	webSocketSender  WebSocketSenderHandler
	isClosed         atomic.Flag
}

// NewWebsocketsDriver will create a new instance of webSocketsDriver
func NewWebsocketsDriver(args ArgsWebSocketsDriver) (*webSocketsDriver, error) {
	if check.IfNil(args.Marshaller) {
		return nil, outportSenderData.ErrNilMarshaller
	}
	if check.IfNil(args.WebsocketSender) {
		return nil, outportSenderData.ErrNilWebSocketSender
	}
	if check.IfNil(args.Log) {
		return nil, outportSenderData.ErrNilLogger
	}

	isClosedFlag := atomic.Flag{}
	isClosedFlag.SetValue(false)

	return &webSocketsDriver{
		marshalizer:     args.Marshaller,
		webSocketSender: args.WebsocketSender,
		log:             args.Log,
		isClosed:        isClosedFlag,
	}, nil
}

// SaveBlock will send the provided block saving arguments within the websocket
func (o *webSocketsDriver) SaveBlock(outportBlock *outport.OutportBlock) error {
	return o.handleAction(outportBlock, outportSenderData.OperationSaveBlock)
}

// RevertIndexedBlock will handle the action of reverting the indexed block
func (o *webSocketsDriver) RevertIndexedBlock(blockData *outport.BlockData) error {
	return o.handleAction(blockData, outportSenderData.OperationRevertIndexedBlock)
}

// SaveRoundsInfo will handle the saving of rounds
func (o *webSocketsDriver) SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error {
	return o.handleAction(roundsInfos, outportSenderData.OperationSaveRoundsInfo)
}

// SaveValidatorsPubKeys will handle the saving of the validators' public keys
func (o *webSocketsDriver) SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error {
	return o.handleAction(validatorsPubKeys, outportSenderData.OperationSaveValidatorsPubKeys)
}

// SaveValidatorsRating will handle the saving of the validators' rating
func (o *webSocketsDriver) SaveValidatorsRating(validatorsRating *outport.ValidatorsRating) error {
	return o.handleAction(validatorsRating, outportSenderData.OperationSaveValidatorsRating)
}

// SaveAccounts will handle the accounts' saving
func (o *webSocketsDriver) SaveAccounts(accounts *outport.Accounts) error {
	return o.handleAction(accounts, outportSenderData.OperationSaveAccounts)
}

// FinalizedBlock will handle the finalized block
func (o *webSocketsDriver) FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error {
	return o.handleAction(finalizedBlock, outportSenderData.OperationFinalizedBlock)
}

// GetMarshaller returns the internal marshaller
func (o *webSocketsDriver) GetMarshaller() marshal.Marshalizer {
	return o.marshalizer
}

func (o *webSocketsDriver) handleAction(args interface{}, operation outportSenderData.OperationType) error {
	if o.isClosed.IsSet() {
		return outportSenderData.ErrWebSocketServerIsClosed
	}

	marshalledPayload, err := o.marshalizer.Marshal(args)
	if err != nil {
		o.log.Error("cannot marshal block", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while marshaling block for operation %s", err, operation.String())
	}

	err = o.webSocketSender.Send(outportSenderData.WsSendArgs{
		Payload: marshalledPayload,
		OpType:  operation,
	})
	if err != nil {
		o.log.Error("cannot send on route", "operation", operation.String(), "error", err)
		return fmt.Errorf("%w while sending data on route for operation %s", err, operation.String())
	}

	return nil
}

// Close will handle the closing of the outport driver web socket sender
func (o *webSocketsDriver) Close() error {
	o.isClosed.SetValue(true)
	return o.webSocketSender.Close()
}

// IsInterfaceNil returns true if there is no value under the interface
func (o *webSocketsDriver) IsInterfaceNil() bool {
	return o == nil
}
