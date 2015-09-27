package game

import (
	"container/heap"
	"fmt"

	"gem/encoding"
	"gem/protocol"
	"gem/runite"
)

type updateService struct {
	runite *runite.Context
	queue  updateQueue
	notify chan int
}

func newUpdateService(runite *runite.Context) *updateService {
	return &updateService{
		runite: runite,
		queue:  make(updateQueue, 0),
		notify: make(chan int, 16),
	}
}

// processQueue resolves requests from the local cache and buffers the responses
// requests are processed in priority order
func (svc *updateService) processQueue() {
	for {
		// block until we're notified by enqueue
		// hopefully more efficient than just spinning in a loop
		<-svc.notify
		for len(svc.queue) > 0 {
			item := heap.Pop(&svc.queue).(*queueItem)
			request := item.request
			conn := item.conn
			log := item.log

			if !conn.active {
				// client is disconnecting. discard
				continue
			}

			data, err := request.Resolve(svc.runite)
			if err != nil {
				log.Errorf(err.Error())
				conn.Disconnect()
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

				response := &protocol.UpdateResponse{
					Index: request.Index,
					File:  request.File,
					Size:  encoding.Int16(len(data)),
					Chunk: encoding.Int8(chunkCount),
					Data:  chunk,
				}

				err := response.Encode(conn, nil)
				if err != nil {
					panic(fmt.Sprintf("unexpected error encoding response: %v", err))
				}

				conn.canWrite <- 1
				wrote += chunkSize
				chunkCount++
			}

		}
	}
}

// enqueue pushes a request into the priority queue for processing
func (svc *updateService) enqueue(conn *GameConnection, request protocol.UpdateRequest) {
	heap.Push(&svc.queue, &queueItem{
		request: request,
		conn:    conn,
		log:     conn.Log.SubModule(request.String()),
	})
	svc.notify <- 1
}

// handleUpdateRequest decodes requests and enqueues them
func (svc *updateService) handleUpdateRequest(ctx *context, b *encoding.Buffer) error {
	var request protocol.UpdateRequest
	if err := request.Decode(b, nil); err != nil {
		return err
	}

	svc.enqueue(ctx.conn, request)
	return nil
}
