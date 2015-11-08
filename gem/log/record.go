package log

import (
	"time"
)

type record struct {
	when time.Time
	lvl  Level
	tag  string
	msg  string
	ctx  Context
}

func (r record) When() time.Time {
	return r.when
}

func (r record) Level() Level {
	return r.lvl
}

func (r record) Tag() string {
	return r.tag
}

func (r record) Message() string {
	return r.msg
}

func (r record) Context() Context {
	return r.ctx
}
