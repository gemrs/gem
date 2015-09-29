import gem

class EventListener(object):
    def register(self):
        gem.event.register_listener(gem.event.Startup, self.startup)
        gem.event.register_listener(gem.event.Shutdown, self.shutdown)
        gem.event.register_listener(gem.event.PreTick, self.pre_tick)
        gem.event.register_listener(gem.event.Tick, self.tick)
        gem.event.register_listener(gem.event.PostTick, self.post_tick)
        gem.event.register_listener(gem.event.PlayerLogin, self.player_login)
        gem.event.register_listener(gem.event.PlayerLogout, self.player_logout)

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

    def player_login(self, event, player):
        pass

    def player_logout(self, event, player):
        pass
