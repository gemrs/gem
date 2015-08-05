import gem
import code
import readline
import signal
import sys

console = code.InteractiveConsole(locals=globals())
logger = gem.syslog.Module("pymain")

def cleanup():
    # Clean up from interact() gracefully..
    console.write("\n")
    logger.Notice("Cleaning up for exit...")
    console.resetbuffer()
    sys.exit(0)

def signal_handler(signal, frame):
    console.write("\n")
    cleanup()

signal.signal(signal.SIGINT, signal_handler)

engine = gem.Engine()
engine.Start()

logger.Info("Finished engine initialization")
logger.Notice("Transferring control to interactive console")

console.interact("")
cleanup()
