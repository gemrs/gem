package protocol_os_161

import "github.com/gemrs/gem/gem/game/entity"

type playerFlags int

const (
	playerFlagIdentityUpdate = (1 << 1)
)

func translatePlayerFlags(flags entity.Flags) playerFlags {
	var newFlags playerFlags
	if flags&entity.MobFlagIdentityUpdate != 0 {
		newFlags |= playerFlagIdentityUpdate
	}
	return newFlags
}
