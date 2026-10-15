class OnlineShoppingSystem
{
    private static volatile OnlineShoppingSystem instance;
    private static readonly object syncRoot = new object();

    private readonly Dictionary<string, Product> products = new Dictionary<string, Product>();
    private readonly Dictionary<string, Customer> customers = new Dictionary<string, Customer>();
    private readonly Dictionary<string, Order> orders = new Dictionary<string, Order>();

    private readonly InventoryService inventoryService;
    private readonly PaymentService paymentService;
    private readonly OrderService orderService;
    private readonly SearchService searchService;

    private OnlineShoppingSystem()
    {
        inventoryService = new InventoryService();
        paymentService = new PaymentService();
        orderService = new OrderService(inventoryService);
        searchService = new SearchService(products.Values);
    }

    public static OnlineShoppingSystem GetInstance()
    {
        if (instance == null)
        {
            lock (syncRoot)
            {
                if (instance == null)
                {
                    instance = new OnlineShoppingSystem();
                }
            }
        }
        return instance;
    }

    public void AddProduct(Product product, int initialStock)
    {
        products[product.GetId()] = product;
        inventoryService.AddStock(product, initialStock);
    }

    public Customer RegisterCustomer(string name, string email, string password, Address address)
    {
        var customer = new Customer(name, email, password, address);
        customers[customer.GetId()] = customer;
        return customer;
    }

    public void AddToCart(string customerId, string productId, int quantity)
    {
        AddToCart(customerId, products[productId], quantity);
    }

    public void AddToCart(string customerId, Product product, int quantity)
    {
        var customer = customers[customerId];
        customer.GetAccount().GetCart().AddItem(product, quantity);
    }

    public ShoppingCart GetCustomerCart(string customerId)
    {
        var customer = customers[customerId];
        return customer.GetAccount().GetCart();
    }

    public List<Product> SearchProducts(string name)
    {
        return searchService.SearchByName(name);
    }

    public Order PlaceOrder(string customerId, IPaymentStrategy paymentStrategy)
    {
        var customer = customers[customerId];
        var cart = customer.GetAccount().GetCart();

        if (!cart.GetItems().Any())
        {
            Console.WriteLine("Cannot place an order with an empty cart.");
            return null;
        }

        // Reserve the complete cart before invoking the existing payment boundary.
        Order order;
        try
        {
            order = orderService.CreateOrder(customer, cart);
        }
        catch (Exception e)
        {
            Console.WriteLine($"Order placement failed: {e.Message}");
            return null;
        }

        bool paymentSuccess = false;
        try
        {
            paymentSuccess = paymentService.ProcessPayment(paymentStrategy, cart.CalculateTotal());
        }
        finally
        {
            if (!paymentSuccess) order.CancelOrder();
        }
        if (!paymentSuccess)
        {
            Console.WriteLine("Payment failed. Please try again.");
            return null;
        }
        orders[order.GetId()] = order;
        cart.ClearCart();
        return order;
    }
}