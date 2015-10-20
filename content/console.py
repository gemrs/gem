import code
import readline
import exceptions

import gem
import gem.event

logger = gem.syslog.Module(__name__)

console = code.InteractiveConsole(locals=globals())

def cleanup(event):
    console.resetbuffer()

def interact():
    gem.Shutdown.Register(gem.event.PyListener(cleanup))
    console.interact("")
