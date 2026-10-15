abstract class Subject
{
    private readonly List<INotificationObserver> observers = new List<INotificationObserver>();
    private readonly object observerLock = new object();

    public void AddObserver(INotificationObserver observer)
    {
        lock (observerLock)
        {
            observers.Add(observer);
        }
    }

    public void RemoveObserver(INotificationObserver observer)
    {
        lock (observerLock)
        {
            observers.Remove(observer);
        }
    }

    public void NotifyObservers(Notification notification)
    {
        INotificationObserver[] observerSnapshot;
        lock (observerLock)
        {
            observerSnapshot = observers.ToArray();
        }
        foreach (var observer in observerSnapshot)
        {
            observer.Update(notification);
        }
    }
}