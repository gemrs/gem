import pytest
import Queue

import gem.event

events = [
    gem.event.TestEvent1,
    gem.event.TestEvent2,
    gem.event.TestEvent3,
    gem.event.TestEvent4,
]

def new_dispatch_test_case(register_events, raise_events, listen_events):
    return {
        'register_events': register_events,
        'raise_events': raise_events,
        'listen_events': listen_events,
    }

dispatch_test_cases = [
    new_dispatch_test_case(
        [events[0]],
        [events[0]],
        [events[0]],
    ),
    new_dispatch_test_case(
        [events[0], events[1]],
        [events[1]],
        [events[1]],
    ),
    new_dispatch_test_case(
        [events[0], events[1]],
        [events[1], events[1], events[0]],
        [events[1], events[1], events[0]],
    ),
    new_dispatch_test_case(
        [events[0]],
        [events[1], events[2], events[3]],
        [],
    ),
    new_dispatch_test_case(
        [events[0], events[1], events[2]],
        [events[1], events[2], events[3]],
        [events[1], events[2]],
    ),
    new_dispatch_test_case(
        [events[0], events[1], events[0]],
        [events[1], events[0]],
        [events[1], events[0], events[0]],
    ),
]

def test_event_dispatch():
    q = Queue.Queue()

    def callback(event):
        q.put(event)

    for tc_id, tc in enumerate(dispatch_test_cases):
        print "Launching test case {0}".format(tc_id)
        gem.event.clear()

        for ev in tc['register_events']:
            gem.event.register_listener(ev, callback)

        for ev in tc['raise_events']:
            gem.event.raise_event(ev)

        for ev in tc['listen_events']:
            event = q.get()
            assert event == ev

def new_args_test_case(register_events, raise_events, listen_events, arg1, arg2):
    return {
        'register_events': register_events,
        'raise_events': raise_events,
        'listen_events': listen_events,
        'arg1': arg1,
        'arg2': arg2,
    }

args_test_cases = [
    new_args_test_case(
        [events[0]],
        [events[0]],
        [events[0]],
        "A_STRING", 1234
    ),
    new_args_test_case(
        [events[0], events[1]],
        [events[1]],
        [events[1]],
        "A_STRING", 1234
    ),
    new_args_test_case(
        [events[0], events[1]],
        [events[1], events[1], events[0]],
        [events[1], events[1], events[0]],
        "A_STRING", 1234
    ),
    new_args_test_case(
        [events[0]],
        [events[1], events[2], events[3]],
        [],
        "A_STRING", 1234
    ),
]

def test_event_dispatch():
    q = Queue.Queue()

    def callback(event, arg1, arg2):
        q.put(event)
        q.put(arg1)
        q.put(arg2)

    for tc_id, tc in enumerate(args_test_cases):
        print "Launching test case {0}".format(tc_id)
        gem.event.clear()

        for ev in tc['register_events']:
            gem.event.register_listener(ev, callback)

        for ev in tc['raise_events']:
            gem.event.raise_event(ev, tc['arg1'], tc['arg2'])

        for ev in tc['listen_events']:
            event = q.get()
            assert event == ev
            assert q.get() == tc['arg1']
            assert q.get() == tc['arg2']
