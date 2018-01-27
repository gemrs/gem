package protocol

const MaxPlayers = 2048

// +gen pack_outgoing
type PlayerUpdate struct {
	Me       PlayerUpdateData
	Others   map[int]PlayerUpdateData
	Removing map[int]bool
	Visible  []int
	Updating []int
	Adding   []int
}
