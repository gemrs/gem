package player

type ClientConfig interface {
	TabInterface(int) int
	SetTabInterface(tab int, id int)
}
