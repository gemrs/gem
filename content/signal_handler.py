import atexit
import os

import gem
import gem.log
import console

logger = gem.log.Module(__name__, None)

def setup_exit_handler(cleanup_func):
    def exit_handler():
        logger.info("Cleaning up for exit...")
        cleanup_func()
        os._exit(0)

    atexit.register(exit_handler)
