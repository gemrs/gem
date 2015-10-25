import pytest

import gem.task
import gem.engine

def test_task_api():
    # For tasks to tick, we need the engine
    engine = gem.engine.Engine()
    engine.Start()

    task_stack = ["pre", "tick", "post"]
    def task(task, userdata):
        assert task_stack.pop(0) == userdata
        if len(task_stack) == 0:
            engine.Stop()
        return False

    scheduler = gem.task.submit(task, gem.task.PreTick, 2, "pre")
    scheduler = gem.task.submit(task, gem.task.Tick, 2, "tick")
    scheduler = gem.task.submit(task, gem.task.PostTick, 2, "post")

    engine.Join()
    assert len(task_stack) == 0
