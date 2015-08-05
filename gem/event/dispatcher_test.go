package event

import (
	"testing"
)

var events = []Event{
	"TEST_EVENT_0",
	"TEST_EVENT_1",
	"TEST_EVENT_2",
	"TEST_EVENT_3",
}

type dispatchTestCase struct {
	register []Event
	raise    []Event
	listen   []Event
}

var dispatchTestCases = []dispatchTestCase{
	dispatchTestCase{
		[]Event{events[0]},
		[]Event{events[0]},
		[]Event{events[0]},
	},
	dispatchTestCase{
		[]Event{events[0], events[1]},
		[]Event{events[1]},
		[]Event{events[1]},
	},
	dispatchTestCase{
		[]Event{events[0], events[1]},
		[]Event{events[1], events[1], events[0]},
		[]Event{events[1], events[1], events[0]},
	},
	dispatchTestCase{
		[]Event{events[0]},
		[]Event{events[1], events[2], events[3]},
		[]Event{},
	},
	dispatchTestCase{
		[]Event{events[0], events[1], events[2]},
		[]Event{events[1], events[2], events[3]},
		[]Event{events[1], events[2]},
	},
	dispatchTestCase{
		[]Event{events[0], events[1], events[0]},
		[]Event{events[1], events[0]},
		[]Event{events[1], events[0], events[0]},
	},
}

func TestDispatch(t *testing.T) {
	for tc_id, tc := range dispatchTestCases {
		t.Logf("Launching test case %v", tc_id)
		events := make(chan Event, 10)

		callback := func(event Event) {
			events <- event
		}

		for _, ev := range tc.register {
			t.Logf("Registering callback for %v", ev)
			Dispatcher.Register(ev, callback)
		}

		for _, ev := range tc.raise {
			t.Logf("Raising event %v", ev)
			Dispatcher.Raise(ev)
		}

		for i, ev := range tc.listen {
			event := <-events
			t.Logf("Caught event %v", ev)
			if event != ev {
				t.Errorf("Listener wasn't called: %v != %v [%v]", event, ev, i)
			}
		}
	}
}
