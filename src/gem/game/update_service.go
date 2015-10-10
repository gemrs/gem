package game

import (
	"gem/encoding"
	"gem/protocol"
	"gem/runite"
)

// updateService represents the internal state of the update servuce
type updateService struct {
	runite *runite.Context
	queue  updateQueue
	notify chan int
}

func newUpdateService(runite *runite.Context) *updateService {
	return &updateService{
		runite: runite,
		queue:  newUpdateQueue(),
		notify: make(chan int, 16),
	}
}

// processQueue resolves requests from the local cache and buffers the responses
// requests are processed in priority order
func (svc *updateService) processQueue() {
	for {
		item := svc.queue.Pop().(*queueItem)
		request := item.request
		conn := item.conn
		log := item.log

		select {
		case <-conn.disconnect:
			// client is disconnecting. discard
			continue
		default:
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

			conn.write <- &protocol.UpdateResponse{
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

// decodeRequest decodes requests and enqueues them
func (svc *updateService) decodeRequest(conn *Connection, b *encoding.Buffer) error {
	var request protocol.UpdateRequest
	if err := request.Decode(b, nil); err != nil {
		return err
	}

	svc.queue.Push(&queueItem{
		request: request,
		conn:    conn,
		log:     conn.Log.SubModule(request.String()),
	})
	return nil
}
