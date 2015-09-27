package game

import (
	"gem/log"
	"gem/protocol"
)

// An queueItem is something we manage in a priority queue.
type queueItem struct {
	request protocol.UpdateRequest
	conn    *GameConnection
	index   int
	log     *log.Module
}

// A updateQueue implements heap.Interface and holds queueItems.
type updateQueue []*queueItem

func (pq updateQueue) Len() int { return len(pq) }

func (pq updateQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].request.Priority > pq[j].request.Priority
}

func (pq updateQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *updateQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*queueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *updateQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
