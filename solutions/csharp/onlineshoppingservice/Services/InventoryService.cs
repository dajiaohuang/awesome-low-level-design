class InventoryService : IOrderObserver
{
    private readonly Dictionary<string, int> stock = new Dictionary<string, int>();
    private readonly object lockObj = new object();
    private readonly HashSet<string> restoredOrders = new HashSet<string>();

    public void Update(Order order)
    {
        lock (lockObj)
        {
            if (order.GetStatus() == OrderStatus.CANCELLED && restoredOrders.Add(order.GetId()))
                foreach (var item in order.GetItems())
                    stock[item.GetProductId()] += item.GetQuantity();
        }
    }

    public void AddStock(Product product, int quantity)
    {
        lock (lockObj)
        {
            if (!stock.ContainsKey(product.GetId()))
            {
                stock[product.GetId()] = 0;
            }
            stock[product.GetId()] += quantity;
        }
    }

    public void UpdateStockForOrder(List<OrderLineItem> items)
    {
        lock (lockObj)
        {
            // Decorated cart lines can share one canonical inventory product.
            var quantities = items.GroupBy(item => item.GetProductId())
                .ToDictionary(group => group.Key, group => group.Sum(item => item.GetQuantity()));
            foreach (var entry in quantities)
            {
                if (!stock.ContainsKey(entry.Key) || stock[entry.Key] < entry.Value)
                {
                    throw new OutOfStockException($"Not enough stock for product ID: {entry.Key}");
                }
            }
            foreach (var entry in quantities)
            {
                stock[entry.Key] -= entry.Value;
            }
        }
    }
}