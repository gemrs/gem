import gem
import inspect

# decorator soup


def listener(cls):
    """listener is a class decorator which adds event handling capabilities"""
    event_hooks = {}

    for name, method in cls.__dict__.iteritems():
        if hasattr(method, "event"):
            # do something with the method and class
            event_hooks[name] = method.event

    orig_init = cls.__init__
    def new_init(self, *args, **kwargs):
        for method, event in event_hooks.iteritems():
            _swap(self, method, event)
        if orig_init:
            orig_init(self, *args, **kwargs)
    cls.__init__ = new_init

    return cls

def callback(event):
    """callback is a method decorator to mark event callbacks.

    To be used within a @listener decorated class."""
    def _callback(fn):
        # fn.event is a tuple of *args and *kwargs for _create_wrapper()
        fn.event = ([event], {})
        return fn
    return _callback

def _swap(obj, method, event):
    """_swap replaces a method with an event decorated wrapper"""
    orig_func = getattr(obj, method)
    wrapper = _create_wrapper(orig_func, event[0][0])
    setattr(obj, method, wrapper)

def _create_wrapper(fn, event, event_passthrough=False):
    """_create_wrapper creates an event decorated wrapper function around a
    method, and registers it as an event callback

    Keyword arguments:
    event_passthrough -- whether the wrapped function expects the event id as its first parameter"""
    args = inspect.getargspec(fn).args
    is_method = bool(args and args[0] == 'self')

    def _wrapper(self, *args, **kwargs):
        if event_passthrough == False:
            # remove the first arg (ignoring self if method)
            if is_method:
                event_idx = 1
            else:
                event_idx = 0
            args = (v for i, v in enumerate(args) if i != event_idx)

        return fn(*args, **kwargs)

    gem.event.register_listener(event, _wrapper)
    return _wrapper
