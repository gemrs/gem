package game

import (
	"gem/encoding"
	"gem/game/server"
	"gem/protocol"
	update_protocol "gem/protocol/update"
	"gem/runite"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type UpdateService -excfield "^[a-z].*" $GOFILE

// updateService represents the internal state of the update servuce
type UpdateService struct {
	py.BaseObject

	runite *runite.Context
	queue  updateQueue
	notify chan int
}

func (svc *UpdateService) Init(runite *runite.Context) error {
	svc.runite = runite
	svc.queue = newUpdateQueue()
	svc.notify = make(chan int, 16)
	go svc.processQueue()
	return nil
}

func (svc *UpdateService) NewClient(conn *server.Connection, service int) server.Client {
	conn.Log().Infof("new update client")
	conn.Write <- new(protocol.OutboundUpdateHandshake)
	return NewUpdateClient(conn, svc)
}

// processQueue resolves requests from the local cache and buffers the responses
// requests are processed in priority order
func (svc *UpdateService) processQueue() {
	for {
		item := svc.queue.Pop().(*queueItem)
		request := item.request
		client := item.client
		log := item.log

		if client.IsDisconnecting() {
			continue
		}

		data, err := request.Resolve(svc.runite)
		if err != nil {
			log.Errorf(err.Error())
			client.Disconnect()
			continue
		}

		wrote := 0
		chunkCount := 0
		for wrote < len(data) {
			chunkSize := 500
			if remaining := len(data) - wrote; remaining < 500 {
				chunkSize = remaining
			}
			chunk := data[wrote : wrote+chunkSize]

			client.Conn().Write <- &update_protocol.OutboundUpdateResponse{
				Index: request.Index,
				File:  request.File,
				Size:  encoding.Int16(len(data)),
				Chunk: encoding.Int8(chunkCount),
				Data:  chunk,
			}

			wrote += chunkSize
			chunkCount++
		}
	}
}
