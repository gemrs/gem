import argparse

import signal_handler
import console
import plugins

import gem
import gem.runite as runite

# Create argparser
parser = argparse.ArgumentParser(description='Gem')
parser.add_argument('--console', action='store_true', help='launch the interactive console')
parser.add_argument('--plugin-path', action='append', help='append to the plugin search path')

logger = gem.syslog.Module("pymain")

def main():
    args = parser.parse_args()

    logger.Notice("Starting Gem v0.9: Opal")

    plugin_path = ["content/plugins"]
    if args.plugin_path is not None:
        plugin_path += args.plugin_path

    plugin_manager = plugins.GemPluginManager(plugin_path)
    plugin_manager.collectPlugins()
    plugin_manager.activatePlugins()

    try:
        gem.runite = runite.Context()
        gem.runite.Unpack("./data/main_file_cache.dat", ["./data/main_file_cache.idx{0}".format(i) for i in range(0, 5)])
    except Exception as e:
        logger.Fatal("Couldn't start unpack game data: {0}".format(e))

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
