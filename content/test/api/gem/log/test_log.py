import pytest

import gem

def test_log_api():
    """Detects changes to the log api"""
    syslog = gem.syslog

    syslog.BeginRedirect()
    syslog.EndRedirect()

    module_log = syslog.Module("test_module")
    log_funcs = ["Critical", "Debug", "Error", "Fatal", "Info", "Notice", "Warning"]

    for fn in log_funcs:
        assert hasattr(module_log, fn)
        func = getattr(module_log, fn)
        func("{0} log message".format(fn))

    submodule_log = module_log.SubModule("test_submodule")

    for fn in log_funcs:
        assert hasattr(submodule_log, fn)
        func = getattr(submodule_log, fn)
        func("{0} log message".format(fn))
