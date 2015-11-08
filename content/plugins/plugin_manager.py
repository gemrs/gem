import logging
from yapsy.PluginManager import PluginManager

import gem
import gem.log

logger = gem.log.Module(__name__, None)

class GemLogHandler(logging.Handler):
    def emit(self, record):
        if record.levelno == logging.CRITICAL:
            logger.error(record.getMessage())
        elif record.levelno == logging.ERROR:
            logger.error(record.getMessage())
        elif record.levelno == logging.WARNING:
            logger.error(record.getMessage())
        elif record.levelno == logging.INFO:
            logger.info(record.getMessage())
        elif record.levelno == logging.DEBUG:
            logger.debug(record.getMessage())

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
                plugin.plugin_object.logger = gem.log.Module(plugin_info.name, None)
            logger.debug("Loading plugin {0}".format(plugin_info.name))
            self.activatePluginByName(plugin_info.name)
