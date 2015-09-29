package event

import (
	"testing"
)

var events = []Event{
	TestEvent1,
	TestEvent2,
	TestEvent3,
	TestEvent4,
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

		Dispatcher.Clear()

		callback := func(event Event, _ ...interface{}) {
			events <- event
		}

		for _, ev := range tc.register {
			Register(ev, callback)
		}

		for _, ev := range tc.raise {
			Raise(ev)
		}

		for i, ev := range tc.listen {
			event := <-events
			if event != ev {
				t.Errorf("Listener wasn't called: %v != %v [%v]", event, ev, i)
			}
		}
	}
}

type argsTestCase struct {
	register []Event
	raise    []Event
	listen   []Event
	arg1     string
	arg2     int
}

var argsTestCases = []argsTestCase{
	{
		[]Event{events[0]},
		[]Event{events[0]},
		[]Event{events[0]},
		"A_STRING",
		1234,
	},
	{
		[]Event{events[0], events[0]},
		[]Event{events[0]},
		[]Event{events[0], events[0]},
		"A_STRING",
		1234,
	},
	{
		[]Event{events[0], events[1]},
		[]Event{events[0]},
		[]Event{events[0]},
		"A_STRING",
		1234,
	},
	{
		[]Event{events[0], events[1]},
		[]Event{events[0], events[1]},
		[]Event{events[0], events[1]},
		"A_STRING",
		1234,
	},
}

func TestArgs(t *testing.T) {
	for tc_id, tc := range argsTestCases {
		t.Logf("Launching test case %v", tc_id)
		evChan := make(chan Event, 10)
		strChan := make(chan string, 10)
		intChan := make(chan int, 10)

		Dispatcher.Clear()

		callback := func(event Event, args ...interface{}) {
			s := args[0].(string)
			i := args[1].(int)
			evChan <- event
			strChan <- s
			intChan <- i
		}

		for _, ev := range tc.register {
			Dispatcher.Register(ev, callback)
		}

		for _, ev := range tc.raise {
			Dispatcher.Raise(ev, tc.arg1, tc.arg2)
		}

		for i, ev := range tc.listen {
			event := <-evChan

			if event != ev {
				t.Errorf("Listener wasn't called: %v != %v [%v]", event, ev, i)
			}

			if <-strChan != tc.arg1 || <-intChan != tc.arg2 {
				t.Errorf("Args incorrect")
			}
		}
	}
}
