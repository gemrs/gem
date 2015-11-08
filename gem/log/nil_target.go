package log

type NilTarget struct{}

func (n NilTarget) Handle(r Record) {}
