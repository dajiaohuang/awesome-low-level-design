class CoffeeVendingMachine
{
    private static CoffeeVendingMachine instance;
    private static readonly object lockObject = new object();
    private IVendingMachineState state;
    private readonly object transactionLock = new object();
    private Coffee selectedCoffee;
    private int moneyInserted;

    private CoffeeVendingMachine()
    {
        state = new ReadyState();
        moneyInserted = 0;
    }

    public static CoffeeVendingMachine GetInstance()
    {
        if (instance == null)
        {
            lock (lockObject)
            {
                if (instance == null)
                {
                    instance = new CoffeeVendingMachine();
                }
            }
        }
        return instance;
    }

    public void SelectCoffee(CoffeeType type, List<ToppingType> toppings)
    {
        lock (transactionLock)
        {
        // 1. Create the base coffee using the factory
        Coffee coffee = CoffeeFactory.CreateCoffee(type);

        // 2. Wrap it with decorators
        foreach (ToppingType topping in toppings)
        {
            switch (topping)
            {
                case ToppingType.EXTRA_SUGAR:
                    coffee = new ExtraSugarDecorator(coffee);
                    break;
                case ToppingType.CARAMEL_SYRUP:
                    coffee = new CaramelSyrupDecorator(coffee);
                    break;
            }
        }
        // Let the state handle the rest
        state.SelectCoffee(this, coffee);
        }
    }

    public void InsertMoney(int amount) { lock (transactionLock) state.InsertMoney(this, amount); }
    public void DispenseCoffee() { lock (transactionLock) state.DispenseCoffee(this); }
    public void Cancel() { lock (transactionLock) state.Cancel(this); }

    // Getters and Setters used by State objects
    public void SetState(IVendingMachineState state) { lock (transactionLock) this.state = state; }
    public IVendingMachineState GetState() { lock (transactionLock) return state; }
    public void SetSelectedCoffee(Coffee selectedCoffee) { lock (transactionLock) this.selectedCoffee = selectedCoffee; }
    public Coffee GetSelectedCoffee() { lock (transactionLock) return selectedCoffee; }
    public void SetMoneyInserted(int moneyInserted) { lock (transactionLock) this.moneyInserted = moneyInserted; }
    public int GetMoneyInserted() { lock (transactionLock) return moneyInserted; }

    public void Reset()
    {
        lock (transactionLock)
        {
            selectedCoffee = null;
            moneyInserted = 0;
        }
    }
}