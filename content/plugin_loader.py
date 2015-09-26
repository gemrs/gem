import plugins

class PluginLoader(object):
    def __init__(self, path):
        self.manager = plugins.GemPluginManager(path)

    def load(self):
        self.manager.collectPlugins()
        self.manager.activatePlugins()
