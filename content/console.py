import code
import readline
import exceptions
import rlcompleter

import gem
import gem.event

logger = gem.syslog.Module(__name__)

console = code.InteractiveConsole(locals=globals())

def cleanup(event):
    console.resetbuffer()

def interact():
    readline.parse_and_bind("tab: complete")
    gem.engine.event.Shutdown.Register(gem.event.PyListener(cleanup))
    console.interact("")
