import pytest
import Queue

import gem.event

events = [
    gem.event.Event("TestEvent1"),
    gem.event.Event("TestEvent2"),
    gem.event.Event("TestEvent3"),
    gem.event.Event("TestEvent4"),
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
        q.queue.clear()
        listeners = {}

        for event in tc['register_events']:
            listener = gem.event.PyListener(callback)
            event.register(listener)
            listeners.setdefault(event, []).append(listener)

        for event in tc['raise_events']:
            event.notify_observers()

        for event in tc['listen_events']:
            raised_event = q.get()
            assert raised_event == event

        for event, listeners in listeners.iteritems():
            for l in listeners:
                event.unregister(l)
            event.notify_observers()
            assert q.empty()

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

def test_event_args():
    q = Queue.Queue()

    def callback(event, arg1, arg2):
        q.put(event)
        q.put(arg1)
        q.put(arg2)

    for tc_id, tc in enumerate(args_test_cases):
        print "Launching test case {0}".format(tc_id)
        q.queue.clear()
        listeners = {}

        for event in tc['register_events']:
            listener = gem.event.PyListener(callback)
            event.register(listener)
            listeners.setdefault(event, []).append(listener)

        for event in tc['raise_events']:
            event.notify_observers(tc['arg1'], tc['arg2'])

        for event in tc['listen_events']:
            raised_event = q.get()
            assert raised_event == event
            assert q.get() == tc['arg1']
            assert q.get() == tc['arg2']

        for event, listeners in listeners.iteritems():
            for l in listeners:
                event.unregister(l)
            event.notify_observers(tc['arg1'], tc['arg2'])
            assert q.empty()

def test_event_dispatch_object():
    q = Queue.Queue()

    class CBClass(object):
        def callback(self, event, arg1, arg2):
            q.put(event)
            q.put(arg1)
            q.put(arg2)

    obj = CBClass()


    for tc_id, tc in enumerate(args_test_cases):
        print "Launching test case {0}".format(tc_id)
        q.queue.clear()
        listeners = {}

        for event in tc['register_events']:
            listener = gem.event.PyListener(obj.callback)
            event.register(listener)
            listeners.setdefault(event, []).append(listener)

        for event in tc['raise_events']:
            event.notify_observers(tc['arg1'], tc['arg2'])

        for event in tc['listen_events']:
            raised_event = q.get()
            assert raised_event == event
            assert q.get() == tc['arg1']
            assert q.get() == tc['arg2']

        for event, listeners in listeners.iteritems():
            for l in listeners:
                event.unregister(l)
            event.notify_observers()
            assert q.empty()
