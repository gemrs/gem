package log_test

import (
	"testing"

	"github.com/gemrs/gem/gem/log"
)

type mockTarget struct {
	records []log.Record
}

func (t *mockTarget) Handle(r log.Record) {
	t.records = append(t.records, r)
}

type entry struct {
	lvl log.Level
	msg string
}

func TestConvenienceFunctions(t *testing.T) {
	log.Targets = map[string]log.Handler{}
	log.Targets["target"] = &mockTarget{make([]log.Record, 0)}

	logger := log.New("testing", log.MapContext{
		"A": 123,
	})

	entries := []entry{
		entry{
			log.LvlInfo,
			"Log Message 1",
		},
		entry{
			log.LvlError,
			"Log Message 2",
		},
		entry{
			log.LvlDebug,
			"Log Message 3",
		},
	}

	logger.Info(entries[0].msg)
	logger.Error(entries[1].msg)
	logger.Debug(entries[2].msg)

	target := log.Targets["target"]
	records := target.(*mockTarget).records

	if len(records) != 3 {
		t.Error("Record length invalid")
	}

	for i, r := range records {
		if r.Level() != entries[i].lvl {
			t.Error("Log level mismatch")
		}

		if r.Tag() != "testing" {
			t.Error("Log tag mismatch")
		}

		if r.Message() != entries[i].msg {
			t.Error("Log message mismatch")
		}

		rctx := r.Context().ContextMap()
		if v, ok := rctx["A"]; !ok || v != 123 {
			t.Error("Log context mismatch")
		}
	}
}

func TestDispatch(t *testing.T) {
	log.Targets = map[string]log.Handler{}
	log.Targets["target1"] = &mockTarget{make([]log.Record, 0)}
	log.Targets["target2"] = &mockTarget{make([]log.Record, 0)}

	logger := log.New("testing", log.MapContext{
		"A": 123,
	})

	entries := []entry{
		entry{
			log.LvlInfo,
			"Log Message 1",
		},
		entry{
			log.LvlDebug,
			"Log Message 2",
		},
	}

	for _, r := range entries {
		logger.Dispatch(r.lvl, r.msg)
	}

	if len(log.Targets) != 2 {
		t.Error("Target length invalid")
	}

	for _, target := range log.Targets {
		records := target.(*mockTarget).records

		if len(records) != 2 {
			t.Error("Record length invalid")
		}

		for i, r := range records {
			if r.Level() != entries[i].lvl {
				t.Error("Log level mismatch")
			}

			if r.Tag() != "testing" {
				t.Error("Log tag mismatch")
			}

			if r.Message() != entries[i].msg {
				t.Error("Log message mismatch")
			}

			rctx := r.Context().ContextMap()
			if v, ok := rctx["A"]; !ok || v != 123 {
				t.Error("Log context mismatch")
			}
		}
	}
}

func TestBuffer(t *testing.T) {
	mtarget := &mockTarget{make([]log.Record, 0)}
	buffer := log.NewBufferingTarget(mtarget)

	log.Targets = map[string]log.Handler{}
	log.Targets["buffer"] = buffer

	logger := log.New("testing", log.MapContext{
		"A": 123,
	})

	entries := []entry{
		entry{
			log.LvlInfo,
			"Log Message 1",
		},
		entry{
			log.LvlDebug,
			"Log Message 2",
		},
	}

	buffer.Redirect()

	logger.Dispatch(entries[0].lvl, entries[0].msg)

	if len(buffer.Buffered()) != 1 {
		t.Error("Buffered length incorrect")
	}

	if len(mtarget.records) != 0 {
		t.Error("Handled length incorrect")
	}

	buffer.Flush()

	if buffer.Buffered() != nil {
		t.Error("Buffer not emptied")
	}

	if len(mtarget.records) != 1 {
		t.Error("Buffer not flushed")
	}

	logger.Dispatch(entries[1].lvl, entries[1].msg)

	if len(mtarget.records) != 2 {
		t.Error("Buffer not flushed")
	}

	records := mtarget.records

	for i, r := range records {
		if r.Level() != entries[i].lvl {
			t.Error("Log level mismatch")
		}

		if r.Tag() != "testing" {
			t.Error("Log tag mismatch")
		}

		if r.Message() != entries[i].msg {
			t.Error("Log message mismatch")
		}

		rctx := r.Context().ContextMap()
		if v, ok := rctx["A"]; !ok || v != 123 {
			t.Error("Log context mismatch")
		}
	}
}
