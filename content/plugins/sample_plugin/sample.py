import gem.task
import plugins.gem_plugin as plugins

class SamplePlugin(plugins.GemPlugin):
    count = 10
    def startup(self, event):
        self.logger.Info("Startup hook")
        gem.task.submit(self.task_test, gem.task.PostTick, 2, None)

    def shutdown(self, event):
        self.logger.Info("Shutdown hook")

    def tick(self, event):
        self.logger.Info("Tick")

    def task_test(self, when, userdata):
        self.logger.Info("Tock.. {0}".format(self.count))
        self.count -= 1
        return self.count > 0
