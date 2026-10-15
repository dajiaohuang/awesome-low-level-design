class PaidState : IVendingMachineState
{
    public void SelectCoffee(CoffeeVendingMachine machine, Coffee coffee)
    {
        Console.WriteLine("Already paid. Please dispense or cancel.");
    }

    public void InsertMoney(CoffeeVendingMachine machine, int amount)
    {
        machine.SetMoneyInserted(machine.GetMoneyInserted() + amount);
        Console.WriteLine($"Additional {amount} inserted. Total: {machine.GetMoneyInserted()}");
    }

    public void DispenseCoffee(CoffeeVendingMachine machine)
    {
        Inventory inventory = Inventory.GetInstance();
        Coffee coffee = machine.GetSelectedCoffee();

        if (!inventory.DeductIngredients(coffee.GetRecipe()))
        {
            Console.WriteLine("Sorry, we are out of ingredients. Refunding your money.");
            Console.WriteLine($"Refunding {machine.GetMoneyInserted()}");
            machine.Reset();
            machine.SetState(new OutOfIngredientState());
            return;
        }

        // Prepare only after the inventory reservation succeeds.
        coffee.Prepare();

        // Calculate change
        int change = machine.GetMoneyInserted() - coffee.GetPrice();
        if (change > 0)
        {
            Console.WriteLine($"Here's your change: {change}");
        }

        machine.Reset();
        machine.SetState(new ReadyState());
    }

    public void Cancel(CoffeeVendingMachine machine)
    {
        Console.WriteLine($"Transaction cancelled. Refunding {machine.GetMoneyInserted()}");
        machine.Reset();
        machine.SetState(new ReadyState());
    }
}