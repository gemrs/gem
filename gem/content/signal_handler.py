import signal
import atexit
import os

import gem
import console

logger = gem.syslog.Module(__name__)

def setup_exit_handler(cleanup_func):
    def exit_handler():
        logger.Notice("Cleaning up for exit...")
        cleanup_func()
        os._exit(0)

    def signal_handler(signal, frame):
        exit_handler()

    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)
    signal.signal(signal.SIGABRT, signal_handler)
    signal.signal(signal.SIGQUIT, signal_handler)
    atexit.register(exit_handler)
