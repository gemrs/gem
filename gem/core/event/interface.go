package event

type Observable interface {
	Key() string
	Register(ObserverIface)
	Unregister(ObserverIface)
	NotifyObservers(...interface{})
}

type ObserverIface interface {
	Id() int
	Notify(*Event, ...interface{})
}
