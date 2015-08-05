package event

type Listener func(Event)

type Listeners []Listener

func (l Listeners) Dispatch(event Event) {
	for _, listener := range l {
		listener(event)
	}
}
