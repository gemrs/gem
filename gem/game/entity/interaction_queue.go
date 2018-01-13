package entity

type Interaction interface {
	Interruptible() bool
	Tick(e Entity) bool
}

type InteractionQueue struct {
	queue []Interaction
}

func NewInteractionQueue() *InteractionQueue {
	return &InteractionQueue{}
}

func (q *InteractionQueue) Append(i Interaction) {
	q.queue = append(q.queue, i)
}

func (q *InteractionQueue) InterruptAndAppend(i Interaction) {
	if len(q.queue) > 0 {
		if q.queue[0].Interruptible() {
			q.Clear()
		}
	}

	q.Append(i)
}

func (q *InteractionQueue) Tick(e Entity) {
	if len(q.queue) == 0 {
		return
	}

	current := q.queue[0]
	completed := current.Tick(e)
	if completed {
		q.queue = q.queue[1:]
	}
}

func (q *InteractionQueue) Clear() {
	q.queue = nil
}
