package entity

import (
	"github.com/sinusoids/gem/gem/util/id"
)

var indexChan <-chan int

func init() {
	indexChan = id.Generator()
}

func NextIndex() int {
	return <-indexChan
}
