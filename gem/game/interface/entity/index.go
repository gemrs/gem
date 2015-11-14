package entity

import (
	"github.com/gemrs/gem/gem/util/id"
)

var indexChan <-chan int

func init() {
	indexChan = id.Generator()
}

func NextIndex() int {
	return <-indexChan
}
