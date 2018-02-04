package data

import (
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/runite"
)

var Huffman *encoding.Huffman

//glua:bind
func LoadHuffmanTable(runite *runite.Context) error {
	idx, err := runite.FS.Index(10)
	if err != nil {
		return err
	}

	huffmanFile := idx.FileIndexByName("huffman")
	container, err := idx.Container(huffmanFile)
	if err != nil {
		return err
	}
	data := container.Data

	Huffman = encoding.NewHuffman(data)

	logger.Notice("Loaded huffman table")
	return nil
}
