class User
{
    private readonly string id;
    private readonly string name;
    private readonly string email;
    private readonly HashSet<User> friends = new HashSet<User>();
    private readonly List<Post> posts = new List<Post>();
    private readonly object collectionLock = new object();

    public User(string name, string email)
    {
        this.id = Guid.NewGuid().ToString();
        this.name = name;
        this.email = email;
    }

    public void AddFriend(User friend)
    {
        lock (collectionLock)
        {
            friends.Add(friend);
        }
    }

    public void AddPost(Post post)
    {
        lock (collectionLock)
        {
            posts.Add(post);
        }
    }

    public string GetId() { return id; }
    public string GetName() { return name; }
    public HashSet<User> GetFriends()
    {
        lock (collectionLock)
        {
            return new HashSet<User>(friends);
        }
    }

    public List<Post> GetPosts()
    {
        lock (collectionLock)
        {
            return new List<Post>(posts);
        }
    }
}