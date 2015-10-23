package rt3

import (
	"bytes"
	"errors"

	"github.com/sinusoids/gem/gem/encoding"
)

var (
	ErrInvalidIndex     = errors.New("index format invalid")
	ErrInvalidData      = errors.New("data format invalid")
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

// todo: derive these constants from the actual structures
const (
	idxSize   int = 6
	blockSize int = 520
)

type buffer []byte

type JagFS struct {
	data    buffer
	indices []*JagFSIndex
}

func UnpackJagFS(data *bytes.Buffer, indices []*bytes.Buffer) (*JagFS, error) {
	var err error
	fs := &JagFS{}
	fs.data = buffer(data.Bytes())
	fs.indices = make([]*JagFSIndex, len(indices))
	for i, index := range indices {
		fs.indices[i], err = unpackFSIndex(fs.data, index)
		if err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func (fs *JagFS) IndexCount() int {
	return len(fs.indices)
}

func (fs *JagFS) Index(index int) (*JagFSIndex, error) {
	if index > len(fs.indices) {
		return nil, ErrIndexOutOfBounds
	}

	return fs.indices[index], nil
}

type JagFSIndex struct {
	fileIndices []FSIndex
	fileCache   map[int]buffer
	data        buffer
}

func unpackFSIndex(data buffer, indexBuffer *bytes.Buffer) (*JagFSIndex, error) {
	if indexBuffer.Len()%idxSize != 0 {
		return nil, ErrInvalidIndex
	}

	index := &JagFSIndex{
		fileIndices: make([]FSIndex, indexBuffer.Len()/idxSize),
		fileCache:   make(map[int]buffer),
		data:        data,
	}

	for i := range index.fileIndices {
		fileIndex := &index.fileIndices[i]
		if err := fileIndex.Decode(indexBuffer, 0); err != nil {
			return nil, ErrInvalidIndex
		}
	}

	return index, nil
}

func (idx *JagFSIndex) FileCount() int {
	return len(idx.fileIndices)
}

func (idx *JagFSIndex) File(index int) ([]byte, error) {
	// we have a cached copy
	if buffer, ok := idx.fileCache[index]; ok {
		return buffer, nil
	}

	// extract the file
	// we can hold 256kb in this channel at any one time.. is this enough?
	if index > len(idx.fileIndices) {
		return nil, ErrIndexOutOfBounds
	}
	fileIndex := idx.fileIndices[index]

	length := int(fileIndex.Length)
	block := int(fileIndex.StartBlock)
	idx.fileCache[index] = make(buffer, 0)

	var err error
	for block != 0 {
		block, err = idx.constructFile(block, length)
		if err != nil {
			return nil, err
		}
		length -= 512
	}

	if length > 0 {
		return nil, ErrInvalidData
	}

	if len(idx.fileCache[index]) != int(fileIndex.Length) {
		return nil, ErrInvalidData
	}

	return idx.fileCache[index], nil
}

func (idx *JagFSIndex) constructFile(blockId, length int) (nextBlockId int, err error) {
	offset := blockId * blockSize
	if offset > len(idx.data) || offset+blockSize > len(idx.data) {
		return 0, ErrIndexOutOfBounds
	}

	var block FSBlock
	block.Data = make(encoding.Bytes, length)
	buffer := bytes.NewBuffer(idx.data[offset : offset+blockSize])
	if err := block.Decode(buffer, 0); err != nil {
		return 0, ErrInvalidData
	}

	index := int(block.FileID)
	nextBlockId = int(block.NextBlock)
	data := []byte(block.Data)
	if length < 512 {
		data = data[:length]
	}
	idx.fileCache[index] = append(idx.fileCache[index], data...)

	return nextBlockId, nil
}
