class Inventory
{
    private static Inventory instance;
    private static readonly object lockObject = new object();
    private readonly Dictionary<Ingredient, int> stock = new Dictionary<Ingredient, int>();
    private readonly object inventoryLock = new object();

    private Inventory() { }

    public static Inventory GetInstance()
    {
        if (instance == null)
        {
            lock (lockObject)
            {
                if (instance == null)
                {
                    instance = new Inventory();
                }
            }
        }
        return instance;
    }

    public void AddStock(Ingredient ingredient, int quantity)
    {
        lock (inventoryLock)
        {
            if (stock.ContainsKey(ingredient))
                stock[ingredient] += quantity;
            else
                stock[ingredient] = quantity;
        }
    }

    public bool HasIngredients(Dictionary<Ingredient, int> recipe)
    {
        lock (inventoryLock)
        {
            return HasIngredientsUnsafe(recipe);
        }
    }

    public bool DeductIngredients(Dictionary<Ingredient, int> recipe)
    {
        lock (inventoryLock)
        {
            if (!HasIngredientsUnsafe(recipe))
            {
                Console.WriteLine("Not enough ingredients to make coffee.");
                return false;
            }
            foreach (var pair in recipe)
            {
                stock[pair.Key] -= pair.Value;
            }
            return true;
        }
    }

    private bool HasIngredientsUnsafe(Dictionary<Ingredient, int> recipe)
    {
        return recipe.All(pair => stock.GetValueOrDefault(pair.Key, 0) >= pair.Value);
    }

    public void PrintInventory()
    {
        lock (inventoryLock)
        {
            Console.WriteLine("--- Current Inventory ---");
            foreach (var pair in stock)
            {
                Console.WriteLine($"{pair.Key}: {pair.Value}");
            }
            Console.WriteLine("-------------------------");
        }
    }
}