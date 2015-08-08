package event

type _Dispatcher struct {
	listeners map[Event]Listeners
}

var Dispatcher = _Dispatcher{
	listeners: make(map[Event]Listeners),
}

// Register a listener on an event
func (d *_Dispatcher) Register(event Event, listener Listener) {
	if _, ok := d.listeners[event]; !ok {
		d.listeners[event] = make(Listeners, 0)
	}

	d.listeners[event] = append(d.listeners[event], listener)
}

// Raise an event, and trigger all listeners of that event
func (d *_Dispatcher) Raise(event Event, args ...interface{}) {
	if _, ok := d.listeners[event]; !ok {
		d.listeners[event] = make(Listeners, 0)
	}

	d.listeners[event].Dispatch(event, args...)
}

// Clears the event:listener map
// Mainly for testing
func (d *_Dispatcher) Clear() {
	d.listeners = make(map[Event]Listeners)
}
