import sys
import argparse

import config
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
args = parser.parse_args()

def main():
    logger.Notice("Starting Gem v0.9: Opal")

    # init
    plugin_init()
    gem.runite = create_runite_context()

    # start the engine
    engine = gem.Engine()
    engine.Start()
    logger.Info("Finished engine initialization")
    signal_handler.setup_exit_handler(engine.Stop)

    # enter interactive console if --console flag is set
    if args.console:
        interactive_console()

    logger.Notice("Press Control-D to toggle the interactive console")
    while True:
        line = sys.stdin.readline()
        if not line: # readline will return "" on EOF
            interactive_console()

def create_runite_context():
    try:
        ctx = runite.Context()
        ctx.Unpack(config.game_data['data_file'], config.game_data['index_files'])
        return ctx
    except Exception as e:
        logger.Fatal("Couldn't start unpack game data: {0}".format(e))

def interactive_console():
    logger.Notice("Transferring control to interactive console")
    gem.syslog.BeginRedirect()
    console.interact()
    gem.syslog.EndRedirect()
    logger.Info("Exited interactive console")

def plugin_init():
    plugin_path = ["content/plugins"]
    if args.plugin_path is not None:
        plugin_path += args.plugin_path

    plugin_manager = plugins.GemPluginManager(plugin_path)
    plugin_manager.collectPlugins()
    plugin_manager.activatePlugins()


if __name__ == "__main__":
    main()
