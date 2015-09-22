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

    atexit.register(exit_handler)
