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

var FullscreenFrame = FrameType{
	Root:        165,
	Window:      6,
	Overlay:     30,
	ChatBox:     1,
	PrivateChat: 4,
	ExpDisplay:  23,
	DataOrbs:    24,
	Tabs: map[protocol.InterfaceTab]int{
		//		protocol.TabInventory:   7,
		protocol.TabAttack:    8,
		protocol.TabSkills:    9,
		protocol.TabQuests:    10,
		protocol.TabInventory: 11,
		protocol.TabEquipment: 12,
		protocol.TabPrayer:    13,
		protocol.TabMagic:     14,
		protocol.TabClanChat:  15,
		protocol.TabFriends:   16,
		protocol.TabIgnore:    17,
		protocol.TabLogout:    18,
		protocol.TabSettings:  19,
		protocol.TabEmotes:    20,
		protocol.TabMusic:     21,
	},
}

var DefaultFrame = FixedFrame

var FixedFrame = FrameType{
	Root:        548,
	Window:      20,
	Overlay:     21,
	ChatBox:     23,
	PrivateChat: 16,
	ExpDisplay:  18,
	DataOrbs:    10,
	Tabs: map[protocol.InterfaceTab]int{
		//		protocol.TabInventory:   63,
		protocol.TabAttack:    65,
		protocol.TabSkills:    66,
		protocol.TabQuests:    67,
		protocol.TabInventory: 68,
		protocol.TabEquipment: 69,
		protocol.TabPrayer:    70,
		protocol.TabMagic:     71,
		protocol.TabClanChat:  72,
		protocol.TabFriends:   73,
		protocol.TabIgnore:    74,
		protocol.TabLogout:    75,
		protocol.TabSettings:  76,
		protocol.TabEmotes:    77,
		protocol.TabMusic:     78,
	},
}

var ResizableFrame = FrameType{
	Root:        161,
	Window:      13,
	Overlay:     14,
	ChatBox:     29,
	PrivateChat: 7,
	ExpDisplay:  9,
	DataOrbs:    28,
	Tabs: map[protocol.InterfaceTab]int{
		//		protocol.TabInventory:   66,
		protocol.TabAttack:    68,
		protocol.TabSkills:    69,
		protocol.TabQuests:    70,
		protocol.TabInventory: 71,
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

var MinimalFrame = FrameType{
	Root:        161,
	Window:      13,
	Overlay:     14,
	ChatBox:     29,
	PrivateChat: 7,
	ExpDisplay:  9,
	DataOrbs:    28,
	Tabs: map[protocol.InterfaceTab]int{
		//		protocol.TabInventory:   66,
		protocol.TabAttack:    68,
		protocol.TabSkills:    69,
		protocol.TabQuests:    70,
		protocol.TabInventory: 71,
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
