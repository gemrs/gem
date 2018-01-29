package protocol_os_157

import (
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

var crcTable = []byte{0x00, 0x00, 0x00, 0x00, 0x88, 0x03, 0x7F, 0xC0, 0xD1, 0x00, 0x00, 0x00, 0x00, 0xCB, 0x2A, 0x66, 0x3C, 0x00, 0x00, 0x00, 0x00, 0xCF, 0x95, 0x5B, 0xD1, 0x00, 0x00, 0x05, 0xB8, 0x59, 0x9D, 0xBE, 0x90, 0x00, 0x00, 0x02, 0x17, 0xD8, 0xC8, 0x4C, 0x55, 0x00, 0x00, 0x00, 0x12, 0xFB, 0xFA, 0xC1, 0x00, 0x00, 0x00, 0x01, 0xDE, 0xDC, 0xCA, 0x4F, 0x63, 0x00, 0x00, 0x00, 0x00, 0xDC, 0x56, 0x8E, 0xC9, 0x00, 0x00, 0x01, 0x87, 0xD3, 0x9C, 0x7E, 0x02, 0x00, 0x00, 0x00, 0x78, 0x21, 0x1B, 0xED, 0x71, 0x00, 0x00, 0x00, 0x00, 0x1E, 0x31, 0x34, 0xA2, 0x00, 0x00, 0x00, 0x00, 0xCA, 0x21, 0x33, 0x15, 0x00, 0x00, 0x00, 0x00, 0x29, 0xED, 0xFF, 0xDC, 0x00, 0x00, 0x02, 0x9D, 0xBD, 0x79, 0xDE, 0x39, 0x00, 0x00, 0x00, 0x02, 0x0D, 0xAC, 0xF8, 0x98, 0x00, 0x00, 0x00, 0x04, 0x8B, 0xF4, 0x00, 0xC2, 0x00, 0x00, 0x00, 0x00, 0x92, 0x7E, 0x6F, 0x21, 0x00, 0x00, 0x00, 0x30}

func (req *InboundUpdateRequest) Resolve(ctx *runite.Context) ([]byte, error) {
	if int(req.Index) == 255 && int(req.File) == 255 {
		return crcTable, nil
	}

	fs := ctx.FS
	indexID := int(req.Index) + 1
	if indexID < 0 || indexID > fs.IndexCount() {
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
	chunkMarker := encoding.Uint8(255)
	chunkSize := 512
	chunk := 0
	fmt.Printf("sending response %#v\n", struc)
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
		//		paddingSize := 0
		if len(data) < thisChunkLen {
			thisChunkLen = len(data)
			//			paddingSize = thisChunkLen - len(data)
		}

		encoding.Bytes(data[:thisChunkLen]).Encode(buf, nil)
		data = data[thisChunkLen:]
		/*
			if paddingSize > 0 {
				padding := make([]byte, paddingSize)
				encoding.Bytes(padding).Encode(buf, nil)
			}*/

		chunk++
	}

	fmt.Printf("encoded file in %v chunks\n", chunk)
}
