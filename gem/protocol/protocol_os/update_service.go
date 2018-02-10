package protocol_os

import (
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/runite"
)

// updateService represents the internal state of the update servuce
//glua:bind
type UpdateService struct {
	runite *runite.Context
	queue  updateQueue
	notify chan int
}

//glua:bind constructor UpdateService
func NewUpdateService(runite *runite.Context) *UpdateService {
	svc := &UpdateService{}
	svc.runite = runite
	svc.queue = newUpdateQueue()
	svc.notify = make(chan int, 16)
	go svc.processQueue()
	return svc
}

func (svc *UpdateService) NewClient(conn *server.Connection, service int) server.GameClient {
	conn.Log().Info("new update client")
	client := NewUpdateClient(conn, svc)
	return client
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
			log.Error(err.Error())
			client.Disconnect()
			continue
		}

		client.Send(&OutboundUpdateResponse{
			Index: request.Index,
			File:  request.File,
			Data:  data,
		})
	}
}
