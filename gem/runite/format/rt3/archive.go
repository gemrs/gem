package rt3

import (
	"io"
	"io/ioutil"

	"github.com/gemrs/gem/gem/core/encoding"
)

type Archive struct {
	Entries  [][]byte
	Capacity int
}

func NewArchive(capacity int) *Archive {
	return &Archive{
		Entries:  make([][]byte, 0),
		Capacity: capacity,
	}
}

func (struc *Archive) Decode(r io.Reader, flags interface{}) {
	archiveData, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	buf := encoding.NewBufferBytes(archiveData)
	buf.Seek(-1, 2)
	chunks := buf.GetU8()
	//chunks = int(archiveData[len(archiveData)-3])
	buf.Seek(0, 0)

	chunkSizes := make([][]int, chunks)
	sizes := make([]int, struc.Capacity)
	buf.Seek((-1)-int64(chunks*struc.Capacity*4), 2)
	for chunk := 0; chunk < chunks; chunk++ {
		chunkSize := 0
		chunkSizes[chunk] = make([]int, struc.Capacity)
		for id := 0; id < struc.Capacity; id++ {
			delta := buf.Get32()
			chunkSize += delta
			chunkSizes[chunk][id] = chunkSize
			sizes[id] += chunkSize
		}
	}

	struc.Entries = make([][]byte, struc.Capacity)
	for id := 0; id < struc.Capacity; id++ {
		struc.Entries[id] = make([]byte, sizes[id])
	}

	buf.Seek(0, 0)
	for chunk := 0; chunk < chunks; chunk++ {
		for id := 0; id < struc.Capacity; id++ {
			chunkSize := chunkSizes[chunk][id]
			chunkData := buf.GetBytes(chunkSize)
			copy(struc.Entries[id], chunkData)
		}
	}
}
