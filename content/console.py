import code
import readline
import exceptions

import gem

logger = gem.syslog.Module(__name__)

console = code.InteractiveConsole(locals=globals())

def cleanup(event):
    console.resetbuffer()

def interact():
    gem.event.register_listener(gem.event.Shutdown, cleanup)
    console.interact("")
