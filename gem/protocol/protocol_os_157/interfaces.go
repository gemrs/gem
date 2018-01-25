package protocol_os_157

import "github.com/gemrs/gem/gem/protocol"

type FrameType struct {
	Root        int
	Window      int
	Overlay     int
	ChatBox     int
	PrivateChat int
	ExpDisplay  int
	DataOrbs    int
	Tabs        map[protocol.InterfaceTab]int
}

var MinimalFrame = FrameType{
	Root:        161,
	Window:      13,
	Overlay:     14,
	ChatBox:     29,
	PrivateChat: 7,
	ExpDisplay:  9,
	DataOrbs:    28,
	Tabs: map[protocol.InterfaceTab]int{
		protocol.TabInventory: 66,
		protocol.TabAttack:    68,
		protocol.TabSkills:    69,
		protocol.TabQuests:    70,
		protocol.TabItemBag:   71,
		protocol.TabEquipment: 72,
		protocol.TabPrayer:    73,
		protocol.TabMagic:     74,
		protocol.TabClanChat:  75,
		protocol.TabFriends:   76,
		protocol.TabIgnore:    77,
		protocol.TabLogout:    78,
		protocol.TabSettings:  79,
		protocol.TabEmotes:    80,
		protocol.TabMusic:     81,
	},
}
