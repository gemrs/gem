import gem
import gem.service.archive as archive
import plugins.gem_plugin as plugins
import config

class ArchiveServer(plugins.GemPlugin):
    started = False

    def startup(self, event):
        try:
            self.server = archive.Server()
            self.server.Start(config.archive_server_listen, gem.runite)
            self.started = True
        except Exception as e:
            self.logger.Critical("Couldn't start archive server: {0}".format(e))

    def shutdown(self, event):
        if self.started == True:
            self.server.Stop()
