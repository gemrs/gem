package event

type Observable interface {
	Key() string
	Register(Observer)
	Unregister(Observer)
	NotifyObservers(...interface{})
}

type Observer interface {
	Id() int
	Notify(*Event, ...interface{})
}
