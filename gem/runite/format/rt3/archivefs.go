package rt3

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"sync"

	"github.com/sinusoids/gem/gem/encoding"
)

var archiveIndices = map[string]int{
	"crc":         0,
	"title":       1,
	"config":      2,
	"interface":   3,
	"media":       4,
	"versionlist": 5,
	"textures":    6,
	"wordenc":     7,
	"sounds":      8,
}

type ArchiveFS struct {
	*JagFSIndex
	crc     bytes.Buffer
	crcOnce sync.Once
}

func NewArchiveFS(i *JagFSIndex) *ArchiveFS {
	return &ArchiveFS{JagFSIndex: i}
}

func (a *ArchiveFS) generateCrc() error {
	file := CRCFile{
		Sum: encoding.Int32(1234),
	}

	for i := 0; i < a.FileCount(); i++ {
		data, err := a.File(i)
		if err != nil {
			return err
		}

		file.Archives[i] = encoding.Int32(crc32.Checksum(data, crc32.IEEETable))
		file.Sum = (file.Sum << 1) + file.Archives[i]
	}

	err := file.Encode(&a.crc, nil)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArchiveFS) ResolveArchive(archive string) (crc []byte, err error) {
	if archive == "crc" {
		a.crcOnce.Do(func() {
			err := a.generateCrc()
			if err != nil {
				// If we can't generate the CRC, then something more serious is wrong. panic
				panic(err)
			}
		})
		return a.crc.Bytes(), nil
	}

	if index, ok := archiveIndices[archive]; ok {
		return a.File(index)
	}

	return nil, fmt.Errorf("no such archive: %v", err)
}
