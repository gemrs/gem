import gem
import gem.log
import gem.task

from abc import ABCMeta, abstractmethod

logger = gem.log.Module(__name__, None)

class Task(object):
    __metaclass__ = ABCMeta

    def __init__(self, interval, when=gem.task.Tick):
        self.interval = interval
        self.when = when

    def submit(self):
        gem.task.submit(self.__execute__, self.when, self.interval, None)

    def __execute__(self, task, userdata):
        return self.execute()

    @abstractmethod
    def execute(self):
        pass
