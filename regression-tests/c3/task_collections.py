import sys
from concurrent.futures import ThreadPoolExecutor
from datetime import date
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[2] / 'solutions/python/taskmanagementsystem'))
from task_management_system import TaskManagementSystem
from task_priority import TaskPriority
from task_status import TaskStatus
from task_sort_strategy import SortByPriority

sys.setswitchinterval(0.000001)
with ThreadPoolExecutor(max_workers=16) as pool:
    instances = list(pool.map(lambda _: TaskManagementSystem.get_instance(), range(100)))
system = instances[0]
assert all(instance is system for instance in instances)
user = system.create_user('Author', 'author@example.com')

def exercise(i):
    task = system.create_task(f'task-{i}', 'body', date(2026, 10, 16), TaskPriority.HIGH, user.id)
    system.list_tasks_by_user(user.id)
    system.list_tasks_by_status(TaskStatus.TODO)
    system.search_tasks('task', SortByPriority())
    system.delete_task(task.get_id())

with ThreadPoolExecutor(max_workers=16) as pool:
    list(pool.map(exercise, range(3000)))
assert system.list_tasks_by_status(TaskStatus.TODO) == []
print('PASS task collection contention')
