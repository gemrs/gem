package game

import (
	"gem/log"
	"gem/protocol"
)

// An queueItem is something we manage in a priority queue.
type queueItem struct {
	request protocol.UpdateRequest
	conn    *GameConnection
	log     *log.Module
}

// A updateQueue is a three-level priority queue
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

func (pq *updateQueue) Push(x interface{}) {
	item := x.(*queueItem)
	priority := item.request.Priority
	if priority < 0 || priority > 3 {
		item.log.Error("Priority out of bounds, dropping request")
		return
	}

	pq.c[priority] <- item
}

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
