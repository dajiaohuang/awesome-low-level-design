class Order : Subject
{
    private readonly string id;
    private readonly Customer customer;
    private readonly List<OrderLineItem> items;
    private readonly Address shippingAddress;
    private readonly double totalAmount;
    private readonly DateTime orderDate;
    private OrderStatus status;
    private IOrderState currentState;
    private readonly object stateLock = new object();

    public Order(Customer customer, List<OrderLineItem> items, Address shippingAddress, double totalAmount)
    {
        this.id = Guid.NewGuid().ToString().Substring(0, 8);
        this.customer = customer;
        this.items = items;
        this.shippingAddress = shippingAddress;
        this.totalAmount = totalAmount;
        this.orderDate = DateTime.Now;
        this.status = OrderStatus.PLACED;
        this.currentState = new PlacedState();
        AddObserver(customer);
    }

    public void ShipOrder() { lock (stateLock) { currentState.Ship(this); } }
    public void DeliverOrder() { lock (stateLock) { currentState.Deliver(this); } }
    public void CancelOrder() { lock (stateLock) { currentState.Cancel(this); } }

    public string GetId() { return id; }
    public OrderStatus GetStatus() { lock (stateLock) { return status; } }
    public void SetState(IOrderState state) { lock (stateLock) { currentState = state; } }
    public void SetStatus(OrderStatus status)
    {
        lock (stateLock)
        {
            this.status = status;
            NotifyObservers(this);
        }
    }
    public List<OrderLineItem> GetItems() { return items; }
}