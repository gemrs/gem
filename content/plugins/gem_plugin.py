from yapsy.IPlugin import IPlugin

import gem

class GemPlugin(IPlugin, EventListener):
    def activate(self):
        # Register event listeners
        super(GemPlugin, self).register()
