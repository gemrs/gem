import code
import readline
import signal
import sys
import argparse

import gem

parser = argparse.ArgumentParser(description='Gem')
parser.add_argument('--console', action='store_true', help='launch the interactive console')

args = parser.parse_args()

console = code.InteractiveConsole(locals=globals())
logger = gem.syslog.Module("pymain")

def cleanup():
    # Clean up from interact() gracefully..
    if args.console:
        console.write("\n")

    logger.Notice("Cleaning up for exit...")

    if args.console:
        console.resetbuffer()

def signal_handler(signal, frame):
    if args.console:
        console.write("\n")

    cleanup()

signal.signal(signal.SIGINT, signal_handler)

engine = gem.Engine()
engine.Start()

logger.Info("Finished engine initialization")

if args.console:
    logger.Notice("Transferring control to interactive console")
    console.interact("")

cleanup()
