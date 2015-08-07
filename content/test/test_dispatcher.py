import pytest
import Queue

import gem.event

events = [
    "TEST_EVENT_0",
    "TEST_EVENT_1",
    "TEST_EVENT_2",
    "TEST_EVENT_3",
]

def new_test_case(register_events, raise_events, listen_events):
    return {
        'register_events': register_events,
        'raise_events': raise_events,
        'listen_events': listen_events,
    }

test_cases = [
    new_test_case(
        [events[0]],
        [events[0]],
        [events[0]],
    ),
    new_test_case(
        [events[0], events[1]],
        [events[1]],
        [events[1]],
    ),
    new_test_case(
        [events[0], events[1]],
        [events[1], events[1], events[0]],
        [events[1], events[1], events[0]],
    ),
    new_test_case(
        [events[0]],
        [events[1], events[2], events[3]],
        [],
    ),
    new_test_case(
        [events[0], events[1], events[2]],
        [events[1], events[2], events[3]],
        [events[1], events[2]],
    ),
    new_test_case(
        [events[0], events[1], events[0]],
        [events[1], events[0]],
        [events[1], events[0], events[0]],
    ),
]

def test_event_dispatch():
    q = Queue.Queue()

    def callback(event):
        q.put(event)

    for tc_id, tc in enumerate(test_cases):
        print "Launching test case {0}".format(tc_id)
        gem.event.clear()

        for ev in tc['register_events']:
            print "Registering callback for {0}".format(ev)
            gem.event.register_listener(ev, callback)

        for ev in tc['raise_events']:
            print "Raising event {0}".format(ev)
            gem.event.raise_event(ev)

        for ev in tc['listen_events']:
            event = q.get()
            print "Caught event {0}".format(ev)
            assert event == ev
