package protocol_os

import (
	"github.com/gemrs/willow/log"
)

// An queueItem is something we manage in a priority queue.
type queueItem struct {
	request  InboundUpdateRequest
	priority int
	client   *UpdateClient
	log      log.Log
}

// A updateQueue is a three-level priority queue, used for prioritising update requests
type updateQueue struct {
	c [3]chan *queueItem
}

func newUpdateQueue() updateQueue {
	return updateQueue{
		c: [3]chan *queueItem{
			make(chan *queueItem, 16),
			make(chan *queueItem, 16),
			make(chan *queueItem, 16),
		},
	}
}

// Push adds a new update request to the queue
func (pq *updateQueue) Push(x interface{}) {
	item := x.(*queueItem)
	priority := item.priority
	if priority < 0 || priority > 3 {
		item.log.Error("Priority out of bounds, dropping request")
		return
	}

	pq.c[priority] <- item
}

// Pop pulls the next update request from the queue in priority order
func (pq *updateQueue) Pop() interface{} {
	for {
		select {
		case item := <-pq.c[0]:
			return item
		default:
			select {
			case item := <-pq.c[1]:
				return item
			default:
				select {
				case item := <-pq.c[2]:
					return item
				default:
				}
			}
		}
	}
}
