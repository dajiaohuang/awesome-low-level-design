class ShoppingCart
{
    private readonly Dictionary<string, CartItem> items = new Dictionary<string, CartItem>();

    public void AddItem(Product product, int quantity)
    {
        var existingItem = items.Values.FirstOrDefault(item =>
            item.GetProduct().GetId() == product.GetId() &&
            item.GetProduct().GetPrice() == product.GetPrice() &&
            item.GetProduct().GetDescription() == product.GetDescription());
        if (existingItem != null)
        {
            existingItem.IncrementQuantity(quantity);
            return;
        }

        string cartKey = items.ContainsKey(product.GetId())
            ? Guid.NewGuid().ToString()
            : product.GetId();
        items[cartKey] = new CartItem(product, quantity);
    }

    public void RemoveItem(string productId)
    {
        foreach (string key in items
                     .Where(entry => entry.Value.GetProduct().GetId() == productId)
                     .Select(entry => entry.Key)
                     .ToList())
        {
            items.Remove(key);
        }
    }

    public Dictionary<string, CartItem> GetItems()
    {
        return new Dictionary<string, CartItem>(items);
    }

    public double CalculateTotal()
    {
        return items.Values.Sum(item => item.GetPrice());
    }

    public void ClearCart()
    {
        items.Clear();
    }
}