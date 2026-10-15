class TaskManagementSystem
{
    private static TaskManagementSystem instance;
    private static readonly object lockObject = new object();
    private readonly object registryLock = new object();
    private readonly Dictionary<string, User> users = new Dictionary<string, User>();
    private readonly Dictionary<string, Task> tasks = new Dictionary<string, Task>();
    private readonly Dictionary<string, TaskList> taskLists = new Dictionary<string, TaskList>();

    private TaskManagementSystem() { }

    public static TaskManagementSystem GetInstance()
    {
        if (instance == null)
        {
            lock (lockObject)
            {
                if (instance == null)
                {
                    instance = new TaskManagementSystem();
                }
            }
        }
        return instance;
    }

    public User CreateUser(string name, string email)
    {
        User user = new User(name, email);
        lock (registryLock) { users[user.GetId()] = user; }
        return user;
    }

    public TaskList CreateTaskList(string listName)
    {
        TaskList taskList = new TaskList(listName);
        lock (registryLock) { taskLists[taskList.GetId()] = taskList; }
        return taskList;
    }

    public Task CreateTask(string title, string description, string dueDate,
                          TaskPriority priority, string createdByUserId)
    {
        User createdBy;
        lock (registryLock)
        {
            if (!users.TryGetValue(createdByUserId, out createdBy))
            {
                throw new ArgumentException("User not found.");
            }
        }

        Task task = new TaskBuilder(title)
                .SetDescription(description)
                .SetDueDate(dueDate)
                .SetPriority(priority)
                .SetCreatedBy(createdBy)
                .Build();

        task.AddObserver(new ActivityLogger());

        lock (registryLock) { tasks[task.GetId()] = task; }
        return task;
    }

    public List<Task> ListTasksByUser(string userId)
    {
        User user;
        List<Task> snapshot;
        lock (registryLock)
        {
            if (!users.TryGetValue(userId, out user))
            {
                return new List<Task>();
            }
            snapshot = tasks.Values.ToList();
        }
        return snapshot.Where(task => task.GetAssignee() == user).ToList();
    }

    public List<Task> ListTasksByStatus(TaskStatus status)
    {
        return SnapshotTasks().Where(task => task.GetStatus() == status).ToList();
    }

    public void DeleteTask(string taskId)
    {
        lock (registryLock) { tasks.Remove(taskId); }
    }

    public List<Task> SearchTasks(string keyword, TaskSortStrategy sortingStrategy)
    {
        List<Task> matchingTasks = new List<Task>();
        foreach (var task in SnapshotTasks())
        {
            if (task.GetTitle().Contains(keyword) || task.GetDescription().Contains(keyword))
            {
                matchingTasks.Add(task);
            }
        }
        sortingStrategy.Sort(matchingTasks);
        return matchingTasks;
    }

    private List<Task> SnapshotTasks()
    {
        lock (registryLock) { return tasks.Values.ToList(); }
    }
}