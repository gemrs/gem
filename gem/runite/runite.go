//glua:bind module gem.runite
package runite

import (
	"bytes"
	"io/ioutil"

	"github.com/gemrs/gem/gem/runite/format/rt3"
)

//go:generate glua .

//glua:bind
type Context struct {
	FS              *rt3.JagFS
	ItemDefinitions []*rt3.ItemDefinition
}

//glua:bind constructor Context
func NewContext() *Context {
	return &Context{}
}

//glua:bind
func (r *Context) Unpack(dataFile string, indexFiles []string) error {
	var err error
	r.FS, err = UnpackJagFSFiles(dataFile, indexFiles)
	if err != nil {
		return err
	}

	r.ItemDefinitions, err = rt3.LoadItemDefinitions(r.FS)
	if err != nil {
		return err
	}

	return nil
}

func UnpackJagFS(data *bytes.Buffer, indices []*bytes.Buffer) (*rt3.JagFS, error) {
	return rt3.UnpackJagFS(data, indices)
}

func UnpackJagFSFiles(dataFile string, indexFiles []string) (*rt3.JagFS, error) {
	var err error
	var dataBuffer *bytes.Buffer
	indexBuffers := make([]*bytes.Buffer, len(indexFiles))

	dataBuffer, err = bufferFile(dataFile)
	if err != nil {
		return nil, err
	}

	for i, f := range indexFiles {
		var buf *bytes.Buffer
		buf, err = bufferFile(f)
		if err != nil {
			return nil, err
		}

		indexBuffers[i] = buf
	}

	return UnpackJagFS(dataBuffer, indexBuffers)
}

func bufferFile(path string) (*bytes.Buffer, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(data), nil
}
