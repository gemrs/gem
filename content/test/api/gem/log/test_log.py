import pytest

import gem.log

def test_log_api():
    """Detects changes to the log api"""

    gem.log.begin_redirect()
    gem.log.end_redirect()

    module_log = gem.log.Module("test_module", None)
    log_funcs = ["debug", "error", "info", "notice"]

    assert hasattr(module_log, "tag")
    assert hasattr(module_log, "ctx")

    for fn in log_funcs:
        assert hasattr(module_log, fn)
        func = getattr(module_log, fn)
        func("{0} log message".format(fn))

    child_log = module_log.child("test_child", None)

    for fn in log_funcs:
        assert hasattr(child_log, fn)
        func = getattr(child_log, fn)
        func("{0} log message".format(fn))
