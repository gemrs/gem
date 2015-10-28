import gem.task
import plugins.gem_plugin as plugins
import task

class SampleTask(task.Task):
    def __init__(self, count, interval=1):
        super(SampleTask, self).__init__(interval)
        self.count = count

    def execute(self):
        logger.info("Tock.. {0}".format(self.count))
        self.count -= 1
        return self.count > 0

class SamplePlugin(plugins.GemPlugin):
    sample_task = SampleTask(10)

    def startup(self, event):
        global logger
        logger = self.logger
        self.logger.info("Startup hook")
        self.sample_task.submit()

    def shutdown(self, event):
        self.logger.info("Shutdown hook")

    def tick(self, event):
        self.logger.info("Tick")
