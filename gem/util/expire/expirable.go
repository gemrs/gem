package expire

type Expirable interface {
	Expired() chan bool
	Expire()
}

// A NonExpirable is an Expirable which cannot expire
type NonExpirable struct {
	expire chan bool
}

func NewNonExpirable() NonExpirable {
	return NonExpirable{make(chan bool)}
}

func (e NonExpirable) Expired() chan bool {
	return e.expire
}

func (e NonExpirable) Expire() {
	// Do nothing
}
