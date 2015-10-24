package player

type BodyPart int

const (
	Torso BodyPart = iota
	Arms
	Legs
	Head
	Hands
	Feet
	Beard
	Hair
	Skin
	BodyPartMax
)

type Anim int

const (
	AnimIdle Anim = iota
	AnimSpotRotate
	AnimWalk
	AnimRotate180
	AnimRotateCCW
	AnimRotateCW
	AnimRun
	AnimMax
)

type Appearance interface {
	Gender() int
	HeadIcon() int
	Model(BodyPart) int
	Color(BodyPart) int
}

type Animations interface {
	Animation(Anim) int
}
