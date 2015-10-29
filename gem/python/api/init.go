// package api ensures that all python apis have been registered and linked
package api

import (
	_ "github.com/sinusoids/gem/gem"
	_ "github.com/sinusoids/gem/gem/archive"
	_ "github.com/sinusoids/gem/gem/auth"
	_ "github.com/sinusoids/gem/gem/engine"
	_ "github.com/sinusoids/gem/gem/engine/event"
	_ "github.com/sinusoids/gem/gem/event"
	_ "github.com/sinusoids/gem/gem/game"
	_ "github.com/sinusoids/gem/gem/game/event"
	_ "github.com/sinusoids/gem/gem/game/player"
	_ "github.com/sinusoids/gem/gem/game/position"
	_ "github.com/sinusoids/gem/gem/game/server"
	_ "github.com/sinusoids/gem/gem/log"
	_ "github.com/sinusoids/gem/gem/runite"
	_ "github.com/sinusoids/gem/gem/task"

	"github.com/sinusoids/gem/gem/python/modules"
)

func init() {
	modules.Link()
}
