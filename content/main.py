import argparse

import gem
import signal_handler
import console

parser = argparse.ArgumentParser(description='Gem')
parser.add_argument('--console', action='store_true', help='launch the interactive console')

logger = gem.syslog.Module("pymain")
engine = gem.Engine()

if __name__ == "__main__":
    args = parser.parse_args()

    engine.Start()
    logger.Info("Finished engine initialization")

    signal_handler.setup_exit_handler(engine.Stop)

    if args.console:
        logger.Notice("Transferring control to interactive console")
        console.interact()
