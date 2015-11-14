// package api ensures that all python apis have been registered and linked
package api

import (
	_ "github.com/gemrs/gem/gem"
	_ "github.com/gemrs/gem/gem/archive"
	_ "github.com/gemrs/gem/gem/auth"
	_ "github.com/gemrs/gem/gem/engine"
	_ "github.com/gemrs/gem/gem/engine/event"
	_ "github.com/gemrs/gem/gem/event"
	_ "github.com/gemrs/gem/gem/game"
	_ "github.com/gemrs/gem/gem/game/event"
	_ "github.com/gemrs/gem/gem/game/player"
	_ "github.com/gemrs/gem/gem/game/position"
	_ "github.com/gemrs/gem/gem/game/server"
	_ "github.com/gemrs/gem/gem/game/world"
	_ "github.com/gemrs/gem/gem/log"
	_ "github.com/gemrs/gem/gem/runite"
	_ "github.com/gemrs/gem/gem/task"

	"github.com/gemrs/gem/gem/python/modules"
)

func init() {
	modules.Link()
}
