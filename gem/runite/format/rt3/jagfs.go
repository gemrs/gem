package rt3

import (
	"bytes"
	"errors"
	"fmt"
	"hash/crc32"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/jzelinskie/whirlpool"
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
	data               buffer
	meta               *JagFSIndex
	indices            []*JagFSIndex
	references         []*ReferenceTable
	checksumTable      ChecksumTable
	checksumTableBytes []byte
}

func UnpackJagFS(data *bytes.Buffer, indices []*bytes.Buffer, meta *bytes.Buffer) (*JagFS, error) {
	var err error
	fs := &JagFS{}
	fs.data = buffer(data.Bytes())

	if meta != nil {
		fs.meta, err = unpackFSIndex(fs.data, meta)
		if err != nil {
			return nil, err
		}
	}

	fs.indices = make([]*JagFSIndex, len(indices))
	for i, index := range indices {
		fs.indices[i], err = unpackFSIndex(fs.data, index)
		if err != nil {
			return nil, err
		}
	}

	fs.references = make([]*ReferenceTable, len(indices))
	for i := range fs.references {
		metaIndex, err := fs.Index(255)
		if err != nil {
			return nil, err
		}

		refTableContainer, err := metaIndex.Container(i)
		if err != nil {
			continue
		}

		fs.references[i] = new(ReferenceTable)
		fs.references[i].Decode(refTableContainer, nil)
	}

	fs.buildChecksumTable()

	return fs, nil
}

func (fs *JagFS) buildChecksumTable() error {
	metaIndex, err := fs.Index(255)
	if err != nil {
		return err
	}

	for i, table := range fs.references {
		rawTable, err := metaIndex.File(i)
		var entry ChecksumEntry
		if err == nil {
			entry.Crc = crc32.ChecksumIEEE(rawTable)

			wp := whirlpool.New()
			wp.Write(rawTable)
			entry.Whirlpool = wp.Sum(nil)

			entry.Version = table.Version
			entry.FileCount = table.Capacity
			entry.Size = table.UncompressedSize()
		}
		fs.checksumTable.AddEntry(entry)
	}

	buf := new(bytes.Buffer)
	fs.checksumTable.Encode(buf, false)

	container := NewContainer(CompressionNone, buf.Bytes())
	buf = new(bytes.Buffer)
	container.Encode(buf, nil)

	fs.checksumTableBytes = buf.Bytes()

	return nil
}

func (fs *JagFS) ChecksumTableBytes() []byte {
	return fs.checksumTableBytes
}

func (fs *JagFS) IndexCount() int {
	return len(fs.indices)
}

func (fs *JagFS) Index(index int) (*JagFSIndex, error) {
	if index == 255 {
		return fs.meta, nil
	}

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
		if err := encoding.TryDecode(fileIndex, indexBuffer, 0); err != nil {
			return nil, ErrInvalidIndex
		}
	}

	return index, nil
}

func (idx *JagFSIndex) FileCount() int {
	return len(idx.fileIndices)
}

func (idx *JagFSIndex) FileReader(index int) (io.Reader, error) {
	data, err := idx.File(index)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.Write(data)
	return &buf, nil
}

func (idx *JagFSIndex) Container(index int) (*Container, error) {
	var container Container
	buf, err := idx.FileReader(index)
	if err != nil {
		return &container, err
	}

	err = encoding.TryDecode(&container, buf, nil)
	return &container, err
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
		block, err = idx.constructFile(index, block, length)
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

func (idx *JagFSIndex) constructFile(index, blockId, length int) (nextBlockId int, err error) {
	offset := blockId * blockSize
	if offset > len(idx.data) || offset+blockSize > len(idx.data) {
		return 0, ErrIndexOutOfBounds
	}

	var block FSBlock
	block.Data = make(encoding.Bytes, length)
	buffer := bytes.NewBuffer(idx.data[offset : offset+blockSize])
	if index > 0xFFFF {
		extBlock := FSBlockExt{
			FSBlock: &block,
		}
		if err := encoding.TryDecode(&extBlock, buffer, 0); err != nil {
			return 0, ErrInvalidData
		}
	} else {
		if err := encoding.TryDecode(&block, buffer, 0); err != nil {
			return 0, ErrInvalidData
		}
	}

	blockIndex := int(block.FileID)
	if blockIndex != index {
		panic(fmt.Errorf("block has unexpected file id: expected %v got %v", index, blockIndex))
	}

	nextBlockId = int(block.NextBlock)
	data := []byte(block.Data)
	if length < 512 {
		data = data[:length]
	}
	idx.fileCache[index] = append(idx.fileCache[index], data...)

	return nextBlockId, nil
}
