package protocol_os_162

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol/protocol_os"
)

const (
	playerFlagIdentityUpdate protocol_os.UpdateFlags = (1 << 1)
	playerFlagChatUpdate     protocol_os.UpdateFlags = (1 << 7)
)

func translatePlayerFlags(flags entity.Flags) protocol_os.UpdateFlags {
	var newFlags protocol_os.UpdateFlags

	if flags&entity.MobFlagIdentityUpdate != 0 {
		newFlags |= playerFlagIdentityUpdate
	}

	if flags&entity.MobFlagChatUpdate != 0 {
		newFlags |= playerFlagChatUpdate
	}

	return newFlags
}
