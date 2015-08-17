import logging
from yapsy.PluginManager import PluginManager

import gem

logger = gem.syslog.Module(__name__)

class GemLogHandler(logging.Handler):
    def emit(self, record):
        logger.Info(record.getMessage())

logging.getLogger("yapsy").addHandler(GemLogHandler())

class GemPluginManager(PluginManager):
    def __init__(self, path):
        super(GemPluginManager, self).__init__()
        super(GemPluginManager, self).setPluginPlaces(path)
        super(GemPluginManager, self).setPluginInfoExtension("plugin")

    def activatePlugins(self):
        for plugin_info in self.getAllPlugins():
            plugin = self.getPluginByName(plugin_info.name)
            if plugin is not None:
                plugin.plugin_object.logger = gem.syslog.Module(plugin_info.name)
            logger.Debug("Loading plugin {0}".format(plugin_info.name))
            self.activatePluginByName(plugin_info.name)
