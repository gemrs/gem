import argparse

import gem
import signal_handler
import console
import plugins

import plugins.sample_plugin.sample

# Create argparser
parser = argparse.ArgumentParser(description='Gem')
parser.add_argument('--console', action='store_true', help='launch the interactive console')
parser.add_argument('--plugin-path', action='append', help='append to the plugin search path')

logger = gem.syslog.Module("pymain")

def main():
    args = parser.parse_args()

    plugin_path = ["content/plugins"]
    if args.plugin_path is not None:
        plugin_path += args.plugin_path

    plugin_manager = plugins.GemPluginManager(plugin_path)
    plugin_manager.collectPlugins()
    plugin_manager.activatePlugins()

    engine = gem.Engine()
    engine.Start()
    logger.Info("Finished engine initialization")

    signal_handler.setup_exit_handler(engine.Stop)

    if args.console:
        logger.Notice("Transferring control to interactive console")
        console.interact()

    # Wait for engine to exit
    engine.Join()

if __name__ == "__main__":
    main()
