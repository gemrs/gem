package protocol_os_163

import "github.com/gemrs/gem/gem/game/entity"

type playerFlags int

const (
	playerFlagIdentityUpdate = (1 << 7)
	playerFlagChatUpdate     = (1 << 1)
)

func translatePlayerFlags(flags entity.Flags) playerFlags {
	var newFlags playerFlags

	if flags&entity.MobFlagIdentityUpdate != 0 {
		newFlags |= playerFlagIdentityUpdate
	}

	if flags&entity.MobFlagChatUpdate != 0 {
		newFlags |= playerFlagChatUpdate
	}

	return newFlags
}
