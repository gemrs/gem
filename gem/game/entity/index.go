package entity

import (
	"github.com/gemrs/gem/gem/util/id"
)

var indexChan <-chan int

func init() {
	indexChan = id.Generator()
	<-indexChan
}

func NextIndex() int {
	return <-indexChan
}
