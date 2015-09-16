from yapsy.IPlugin import IPlugin

import gem

class GemPlugin(IPlugin):
    def activate(self):
        gem.event.register_listener(gem.event.Startup, self.startup)
        gem.event.register_listener(gem.event.Shutdown, self.shutdown)
        gem.event.register_listener(gem.event.PreTick, self.pre_tick)
        gem.event.register_listener(gem.event.Tick, self.tick)
        gem.event.register_listener(gem.event.PostTick, self.post_tick)

    def startup(self, event):
        pass

    def shutdown(self, event):
        pass

    def pre_tick(self, event):
        pass

    def tick(self, event):
        pass

    def post_tick(self, event):
        pass
