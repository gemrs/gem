package protocol_os_163

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/runite"
)

const (
	updateFileRequest int = iota
	updateFileRequestPrio
	updateClientLogIn
	updateClientLogOut
	updateEncKeys
	updateClientConnected
	updateClientDisconnected
)

type InboundUpdateHeader struct {
	Id int
}

func (struc *InboundUpdateHeader) Decode(buf io.Reader, flags interface{}) {
	var tmp8 encoding.Uint8
	tmp8.Decode(buf, encoding.IntNilFlag)
	struc.Id = int(tmp8)
}

type InboundConnectionStatus struct {
	Status encoding.Int24
}

func (struc *InboundConnectionStatus) Decode(buf io.Reader, flags interface{}) {
	struc.Status.Decode(buf, encoding.IntNilFlag)
}

type InboundXorKey struct {
	Key encoding.Uint8
	Unk encoding.Uint16
}

func (struc *InboundXorKey) Decode(buf io.Reader, flags interface{}) {
	struc.Key.Decode(buf, encoding.IntNilFlag)
	struc.Unk.Decode(buf, encoding.IntNilFlag)
}

func (struc *InboundXorKey) Encode(buf io.Writer, flags interface{}) {
	body := flags.(encoding.Encodable)

	if int(struc.Key) == 0 {
		body.Encode(buf, nil)
		return
	}

	var plaintext bytes.Buffer
	body.Encode(&plaintext, nil)
	ciphertext := make([]byte, plaintext.Len())
	for i, b := range plaintext.Bytes() {
		ciphertext[i] = b ^ byte(struc.Key)
	}
	buf.Write(ciphertext)
}

type InboundUpdateRequest struct {
	Index encoding.Uint8
	File  encoding.Uint16
}

func (struc *InboundUpdateRequest) Encode(buf io.Writer, flags interface{}) {
	struc.Index.Encode(buf, encoding.IntNilFlag)
	struc.File.Encode(buf, encoding.IntNilFlag)
}

func (struc *InboundUpdateRequest) Decode(buf io.Reader, flags interface{}) {
	struc.Index.Decode(buf, encoding.IntNilFlag)
	struc.File.Decode(buf, encoding.IntNilFlag)
}

func (req *InboundUpdateRequest) String() string {
	return fmt.Sprintf("Cache %v, File %v", req.Index, req.File)
}

func (req *InboundUpdateRequest) Resolve(ctx *runite.Context) ([]byte, error) {
	if int(req.Index) == 255 && int(req.File) == 255 {
		return ctx.FS.ChecksumTableBytes(), nil
	}

	fs := ctx.FS
	indexID := int(req.Index)
	if indexID < 0 || (indexID > fs.IndexCount() && indexID != 255) {
		return nil, fmt.Errorf("cache index out of bounds")
	}

	index, err := fs.Index(indexID)
	if err != nil {
		return nil, fmt.Errorf("error accessing index: %v", err)
	}

	if req.File < 0 || int(req.File) > index.FileCount() {
		return nil, fmt.Errorf("file index out of bounds")
	}

	data, err := index.File(int(req.File))
	if err != nil {
		return nil, fmt.Errorf("error accessing file: %v", err)
	}

	return data, err
}

type OutboundUpdateResponse struct {
	Index encoding.Uint8
	File  encoding.Uint16
	Data  []byte
}

func (res *OutboundUpdateResponse) String() string {
	return fmt.Sprintf("Cache %v, File %v", res.Index, res.File)
}

func (struc *OutboundUpdateResponse) Encode(buf io.Writer, flags interface{}) {
	data := struc.Data
	if struc.Index != 255 {
		// Trim the version
		data = data[:len(data)-2]
	}
	chunkMarker := encoding.Uint8(255)
	chunkSize := 512
	chunk := 0
	for len(data) > 0 {
		headerSize := 0
		if chunk == 0 {
			struc.Index.Encode(buf, encoding.IntNilFlag)
			struc.File.Encode(buf, encoding.IntNilFlag)
			headerSize = 3
		} else {
			chunkMarker.Encode(buf, encoding.IntNilFlag)
			headerSize = 1
		}

		thisChunkLen := chunkSize - headerSize
		if len(data) < thisChunkLen {
			thisChunkLen = len(data)
		}

		encoding.Bytes(data[:thisChunkLen]).Encode(buf, nil)
		data = data[thisChunkLen:]
		chunk++
	}
}
