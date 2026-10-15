abstract class CommentableEntity
{
    protected readonly string id;
    protected readonly User author;
    protected readonly string content;
    protected readonly DateTime timestamp;
    private readonly HashSet<User> likes = new HashSet<User>();
    protected readonly List<Comment> comments = new List<Comment>();
    private readonly object collectionLock = new object();

    public CommentableEntity(User author, string content)
    {
        this.id = Guid.NewGuid().ToString();
        this.author = author;
        this.content = content;
        this.timestamp = DateTime.Now;
    }

    public void AddLike(User user)
    {
        lock (collectionLock)
        {
            likes.Add(user);
        }
    }

    public void AddComment(Comment comment)
    {
        lock (collectionLock)
        {
            comments.Add(comment);
        }
    }

    public string GetId() { return id; }
    public User GetAuthor() { return author; }
    public string GetContent() { return content; }
    public DateTime GetTimestamp() { return timestamp; }
    public List<Comment> GetComments()
    {
        lock (collectionLock)
        {
            return new List<Comment>(comments);
        }
    }

    public HashSet<User> GetLikes()
    {
        lock (collectionLock)
        {
            return new HashSet<User>(likes);
        }
    }
}