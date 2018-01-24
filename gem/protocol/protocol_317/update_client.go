package protocol_317

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol/update_protocol"
	"github.com/gemrs/willow/log"
)

// UpdateClient is a client which serves update requests
type UpdateClient struct {
	*server.Connection
	service *UpdateService
}

// NewUpdateClient constructs a new UpdateClient
func NewUpdateClient(conn *server.Connection, svc *UpdateService) *UpdateClient {
	return &UpdateClient{
		Connection: conn,
		service:    svc,
	}
}

// Conn returns the underlying Connection
func (client *UpdateClient) Conn() *server.Connection {
	return client.Connection
}

// Decode processes incoming requests and adds them to the request queue
func (client *UpdateClient) Decode() error {
	var request update_protocol.InboundUpdateRequest
	err := encoding.TryDecode(&request, client.Conn().ReadBuffer, nil)
	if err != nil {
		return err
	}

	client.service.queue.Push(&queueItem{
		request: request,
		client:  client,
		log:     client.Log().Child("request", log.MapContext{"request": request.String()}),
	})
	return nil
}

// Encode writes encoding.Encodables to the client's buffer
func (client *UpdateClient) Encode(codable encoding.Encodable) error {
	return encoding.TryEncode(codable, client.Conn().WriteBuffer, nil)
}

func (client *UpdateClient) SetDecodeFunc(d server.DecodeFunc) {
	panic("SetDecodeFunc not implemented for UpdateClient")
}

func (client *UpdateClient) IsaacIn() *isaac.ISAAC {
	panic("IsaacIn not implemented for UpdateClient")
}
