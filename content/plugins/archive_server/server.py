import gem
import gem.service.archive as archive
import plugins.gem_plugin as plugins
import config

class ArchiveServer(plugins.GemPlugin):
    def startup(self, event):
        global logger
        logger = self.logger
        try:
            self.server = archive.Server()
            self.server.Start(config.archive_server_listen, gem.runite)
        except Exception as e:
            logger.Critical("Couldn't start archive server: {0}".format(e))

    def shutdown(self, event):
        self.server.Stop()
