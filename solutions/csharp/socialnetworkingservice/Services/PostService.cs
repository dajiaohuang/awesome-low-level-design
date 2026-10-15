class PostService
{
    private readonly PostRepository postRepository = PostRepository.GetInstance();
    private readonly List<IPostObserver> observers = new List<IPostObserver>();
    private readonly object observerLock = new object();

    public void AddObserver(IPostObserver observer)
    {
        lock (observerLock)
        {
            observers.Add(observer);
        }
    }

    public Post CreatePost(User author, string content)
    {
        Post post = new Post(author, content);
        postRepository.Save(post);
        author.AddPost(post);
        IPostObserver[] observerSnapshot;
        lock (observerLock)
        {
            observerSnapshot = observers.ToArray();
        }
        foreach (var observer in observerSnapshot)
        {
            observer.OnPostCreated(post);
        }
        return post;
    }

    public void LikePost(User user, string postId)
    {
        Post post = postRepository.FindById(postId);
        post.AddLike(user);
        IPostObserver[] observerSnapshot;
        lock (observerLock)
        {
            observerSnapshot = observers.ToArray();
        }
        foreach (var observer in observerSnapshot)
        {
            observer.OnLike(post, user);
        }
    }

    public void AddComment(User author, string commentableId, string content)
    {
        Comment comment = new Comment(author, content);
        Post post = postRepository.FindById(commentableId);
        post.AddComment(comment);
        IPostObserver[] observerSnapshot;
        lock (observerLock)
        {
            observerSnapshot = observers.ToArray();
        }
        foreach (var observer in observerSnapshot)
        {
            observer.OnComment(post, comment);
        }
    }
}