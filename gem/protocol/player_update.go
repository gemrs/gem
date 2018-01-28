package protocol

const MaxPlayers = 2048

// +gen pack_outgoing
type PlayerUpdate struct {
	Me       Player
	Others   map[int]Player
	Removing map[int]bool
	Visible  []int
	Updating []int
	Adding   []int
}
