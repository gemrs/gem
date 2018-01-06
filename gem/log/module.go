package log

import (
	"github.com/gemrs/willow/log"
)

//go:generate glua .

//glua:bind
type Module struct {
	log.Log
}

//glua:bind constructor Module
func NewModule(tag string) *Module {
	// TODO context
	return &Module{
		Log: log.New(tag, nil),
	}
}

//glua:bind
func (m *Module) Debug(message string) {
	m.Log.Debug(message)
}

//glua:bind
func (m *Module) Error(message string) {
	m.Log.Error(message)
}

//glua:bind
func (m *Module) Info(message string) {
	m.Log.Info(message)
}

//glua:bind
func (m *Module) Notice(message string) {
	m.Log.Notice(message)
}
