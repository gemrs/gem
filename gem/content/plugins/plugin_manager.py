import logging
from yapsy.PluginManager import PluginManager

import gem

logger = gem.syslog.Module(__name__)

class GemLogHandler(logging.Handler):
    def emit(self, record):
        if record.levelno == logging.CRITICAL:
            logger.Critical(record.getMessage())
        elif record.levelno == logging.ERROR:
            logger.Error(record.getMessage())
        elif record.levelno == logging.WARNING:
            logger.Warning(record.getMessage())
        elif record.levelno == logging.INFO:
            logger.Info(record.getMessage())
        elif record.levelno == logging.DEBUG:
            logger.Debug(record.getMessage())

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
