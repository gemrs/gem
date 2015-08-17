import plugins.gem_plugin as plugins

class SamplePlugin(plugins.GemPlugin):
    def startup(self, event):
        self.logger.Info("Startup hook")

    def shutdown(self, event):
        self.logger.Info("Shutdown hook")

    def tick(self, event):
        self.logger.Info("Tick")
