package runite

type JagFS interface {
	IndexCount() int
	Index(index int) (JagFSIndex, error)
}

type JagFSIndex interface {
	FileCount() int
	File(index int) ([]byte, error)
}
