package event

type _Dispatcher struct {
	listeners map[Event]Listeners
}

var Dispatcher = _Dispatcher{
	listeners: make(map[Event]Listeners),
}

func (d *_Dispatcher) Register(event Event, listener Listener) {
	if _, ok := d.listeners[event]; !ok {
		d.listeners[event] = make(Listeners, 0)
	}

	d.listeners[event] = append(d.listeners[event], listener)
}

func (d *_Dispatcher) Raise(event Event) {
	if _, ok := d.listeners[event]; !ok {
		d.listeners[event] = make(Listeners, 0)
	}

	d.listeners[event].Dispatch(event)
}
