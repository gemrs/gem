package log

import (
	"fmt"
	"strings"
	"time"
)

// MockModule is a logger which tests can use in place of a real logger.
// Provides methods to help test log output.
type MockModule struct {
	tag    string
	ctx    Context
	buffer *BufferingTarget
}

func NewMock(tag string, ctx Context) *MockModule {
	module := &MockModule{
		tag:    tag,
		ctx:    ctx,
		buffer: NewBufferingTarget(NilTarget{}),
	}
	module.buffer.Redirect()
	return module
}

func (m *MockModule) HasLogged(msg string) bool {
	for _, record := range m.buffer.Buffered() {
		if strings.Contains(record.Message(), msg) {
			return true
		}
	}

	return false
}

func (m *MockModule) Dispatch(lvl Level, msg string) {
	record := record{
		when: time.Now(),
		lvl:  lvl,
		tag:  m.tag,
		msg:  msg,
		ctx:  m.ctx,
	}

	m.buffer.Handle(record)
}

func (m *MockModule) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlInfo, msg)
}

func (m *MockModule) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlError, msg)
}

func (m *MockModule) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlDebug, msg)
}

func (m *MockModule) Notice(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m.Dispatch(LvlNotice, msg)
}

func (m *MockModule) Child(tag string, ctx Context) Log {
	if ctx == nil {
		ctx = m.ctx
	}
	return New(m.tag+"/"+tag, ctx)
}
