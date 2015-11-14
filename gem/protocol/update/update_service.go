package update

import (
	"fmt"

	"github.com/gemrs/gem/gem/runite"
)

//go:generate bbc update_service.bb update_service.bb.go

func (req *InboundUpdateRequest) String() string {
	return fmt.Sprintf("Cache %v, File %v, Priority %v", req.Index, req.File, req.Priority)
}

func (res *OutboundUpdateResponse) String() string {
	return fmt.Sprintf("Cache %v, File %v, Size %v, Chunk %v", res.Index, res.File, res.Size, res.Chunk)
}

func (req *InboundUpdateRequest) Resolve(ctx *runite.Context) ([]byte, error) {
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
