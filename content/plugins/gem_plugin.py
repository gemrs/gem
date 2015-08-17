from yapsy.IPlugin import IPlugin

import gem

class GemPlugin(IPlugin):
    def activate(self):
        gem.event.register_listener(gem.event.Startup, self.startup)
        gem.event.register_listener(gem.event.Shutdown, self.shutdown)

    def startup(self, event):
        pass

    def shutdown(self, event):
        pass
