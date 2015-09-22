import gem.service.archive as archive
import plugins.gem_plugin as plugins

class ArchiveServer(plugins.GemPlugin):
    def startup(self, event):
        global logger
        logger = self.logger
        try:
            self.server = archive.Server()
            self.server.Start(":43595")
        except Exception as e:
            logger.Critical("Couldn't start archive server: {0}".format(e))

    def shutdown(self, event):
        self.server.Stop()
