package player

var RevisionConstants = struct {
	InventoryInterfaceId int
}{}

//glua:bind
func SetInventoryInterfaceId(i int) {
	RevisionConstants.InventoryInterfaceId = i
}
