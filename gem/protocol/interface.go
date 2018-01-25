package protocol

//glua:bind
type InterfaceTab int

//glua:bind
const (
	TabAttack InterfaceTab = iota
	TabSkills
	TabQuests
	TabInventory
	TabEquipment
	TabPrayer
	TabMagic
	TabUnused
	TabFriends
	TabIgnore
	TabLogout
	TabSettings
	TabRun
	TabMusic

	TabItemBag
	TabClanChat
	TabEmotes
)

type FrameType interface{}
