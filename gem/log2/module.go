package log

import (
	"fmt"
	"time"
)

var Targets = map[string]Handler{}

type Module struct {
	tag     string
	ctx     Context
	targets []Handler
}

func New(tag string, ctx Context) Log {
	return &Module{
		tag: tag,
		ctx: ctx,
	}
}

func (m *Module) Dispatch(lvl Level, msg string) {
	record := record{
		when: time.Now(),
		lvl:  lvl,
		tag:  m.tag,
		msg:  msg,
		ctx:  m.ctx,
	}

	for _, handler := range Targets {
		handler.Handle(record)
	}
}

func (m *Module) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlInfo, msg)
}

func (m *Module) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlError, msg)
}

func (m *Module) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlDebug, msg)
}
