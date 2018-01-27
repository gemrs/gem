package protocol_os_157

type playerData struct {
	skipFlags [2048]int
	frame     FrameType
}

func newPlayerData() *playerData {
	return &playerData{
		frame: FixedFrame,
	}
}

func getPlayerData(d interface{}) *playerData {
	return d.(*playerData)
}
