package protocol_os

import (
	"bytes"
	"fmt"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/willow/log"
)

// UpdateClient is a client which serves update requests
type UpdateClient struct {
	*server.Connection
	service *UpdateService
	xorKey  InboundXorKey
}

// NewUpdateClient constructs a new UpdateClient
func NewUpdateClient(conn *server.Connection, svc *UpdateService) *UpdateClient {
	return &UpdateClient{
		Connection: conn,
		service:    svc,
	}
}

func (client *UpdateClient) SetProtoData(d interface{}) {
	panic("SetProtoData not implemented")
}

// Conn returns the underlying Connection
func (client *UpdateClient) Conn() *server.Connection {
	return client.Connection
}

// Decode processes incoming requests and adds them to the request queue
func (client *UpdateClient) Decode() error {
	var header InboundUpdateHeader
	header.Decode(client.Conn().ReadBuffer, nil)

	var priority bool
	switch header.Id {
	case updateFileRequestPrio:
		priority = true
		fallthrough
	case updateFileRequest:
		var request InboundUpdateRequest
		request.Decode(client.Conn().ReadBuffer, nil)
		return client.serveFileRequest(request, priority)

	case updateClientLogIn:
		fallthrough
	case updateClientLogOut:
		fallthrough
	case updateClientConnected:
		fallthrough
	case updateClientDisconnected:
		var status InboundConnectionStatus
		status.Decode(client.Conn().ReadBuffer, nil)
		return client.handleConnectionStatus(header.Id, status)

	case updateEncKeys:
		var key InboundXorKey
		key.Decode(client.Conn().ReadBuffer, nil)
		return client.handleKeyUpdate(key)

	default:
		panic(fmt.Errorf("unknown update packet"))
	}
}

func (client *UpdateClient) handleConnectionStatus(id int, status InboundConnectionStatus) error {
	if int(status.Status) != 0 {
		client.Conn().Disconnect()
	}
	return nil
}

func (client *UpdateClient) handleKeyUpdate(key InboundXorKey) error {
	client.xorKey = key
	return nil
}

func (client *UpdateClient) serveFileRequest(request InboundUpdateRequest, prioritize bool) error {
	priority := 1
	if prioritize {
		priority = 0
	}

	client.service.queue.Push(&queueItem{
		request:  request,
		priority: priority,
		client:   client,
		log:      client.Log().Child("request", log.MapContext{"request": request.String()}),
	})
	return nil
}

// Send writes encoding.Encodables to the clients's buffer
func (client *UpdateClient) Send(codable encoding.Encodable) error {
	var buf bytes.Buffer
	err := encoding.TryEncode(&client.xorKey, &buf, codable)
	if err != nil {
		return err
	}
	client.Write <- encoding.Encoded{buf}
	return nil
}

func (client *UpdateClient) SetDecodeFunc(d server.DecodeFunc) {
	panic("SetDecodeFunc not implemented for UpdateClient")
}

func (client *UpdateClient) IsaacIn() *isaac.ISAAC {
	panic("IsaacIn not implemented for UpdateClient")
}
