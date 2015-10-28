import pytest

import gem

def test_log_api():
    """Detects changes to the log api"""
    syslog = gem.syslog

    syslog.begin_redirect()
    syslog.end_redirect()

    module_log = syslog.module("test_module")
    log_funcs = ["critical", "debug", "error", "info", "notice", "warning"]

    for fn in log_funcs:
        assert hasattr(module_log, fn)
        func = getattr(module_log, fn)
        func("{0} log message".format(fn))

    submodule_log = module_log.submodule("test_submodule")

    for fn in log_funcs:
        assert hasattr(submodule_log, fn)
        func = getattr(submodule_log, fn)
        func("{0} log message".format(fn))
