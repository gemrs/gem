package runite

import (
	"io/ioutil"
	"bytes"

	"gem/runite/format/rt3"
)

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
