package server

import (
	"errors"
	"sync"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func createArgs() ArgsWebSocketsServer {
	return ArgsWebSocketsServer{
		RetryDurationInSeconds:   1,
		BlockingAckOnError:       false,
		WithAcknowledge:          false,
		URL:                      "url",
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
	}
}

func TestNewWebSocketsServer(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		args := createArgs()
		ws, err := NewWebSocketsServer(args)
		require.NotNil(t, ws)
		require.Nil(t, err)
	})

	t.Run("empty url, should return error", func(t *testing.T) {
		args := createArgs()
		args.URL = ""
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrEmptyUrl)
	})

	t.Run("nil uint64 byte slice converter, should return error", func(t *testing.T) {
		args := createArgs()
		args.Uint64ByteSliceConverter = nil
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrNilUint64ByteSliceConverter)
	})

	t.Run("zero retry duration in seconds, should return error", func(t *testing.T) {
		args := createArgs()
		args.RetryDurationInSeconds = 0
		ws, err := NewWebSocketsServer(args)
		require.Nil(t, ws)
		require.Equal(t, err, data.ErrZeroValueRetryDuration)
	})
}

func TestServer_ListenAndClose(t *testing.T) {
	args := createArgs()
	args.URL = "localhost:9211"
	wsServer, _ := NewWebSocketsServer(args)

	called := false
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wsServer.Listen()
		wg.Done()
		called = true
	}()

	_ = wsServer.Close()
	wg.Wait()
	require.True(t, called)
}

func TestServer_ListenAndRegisterPayloadHandlerAndClose(t *testing.T) {
	args := createArgs()
	args.URL = "localhost:9211"
	wsServer, _ := NewWebSocketsServer(args)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wsServer.Listen()
		wg.Done()
	}()

	wsServer.RegisterPayloadHandler(&testscommon.PayloadHandlerStub{})
	wsServer.connectionHandler(&testscommon.WebsocketConnectionStub{
		ReadMessageCalled: func() (messageType int, payload []byte, err error) {
			return 0, nil, errors.New("local error")
		},
	})

	_ = wsServer.Close()
	wg.Wait()
}
