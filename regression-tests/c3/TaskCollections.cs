using System;
using System.Linq;
using System.Threading.Tasks;

class TaskCollectionsTest
{
    public static void Main()
    {
        var system = TaskManagementSystem.GetInstance();
        var user = system.CreateUser("Author", "author@example.com");
        Parallel.For(0, 2000, i =>
        {
            var task = system.CreateTask("task-" + i, "body", "2026-10-16", TaskPriority.HIGH, user.GetId());
            system.ListTasksByUser(user.GetId());
            system.ListTasksByStatus(TaskStatus.TODO);
            system.SearchTasks("task", new SortByPriority());
            system.DeleteTask(task.GetId());
        });
        if (system.ListTasksByStatus(TaskStatus.TODO).Count != 0) throw new Exception("Deleted tasks remain");
        Console.WriteLine("PASS task collection contention");
    }
}
