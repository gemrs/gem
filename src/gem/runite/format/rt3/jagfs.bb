type FSIndex struct {
    Length     int24
    StartBlock int24
}

type FSBlock struct {
	FileID       int16
	FilePosition int16
	NextBlock    int24
	Partition    int8
	Data         byte[512]
}
