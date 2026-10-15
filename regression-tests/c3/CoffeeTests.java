package coffeevendingmachine;

/** Runs the tracked assertions when a JUnit launcher is not available. */
public class CoffeeTests {
    public static void main(String[] args) {
        CoffeeVendingMachineTest test = new CoffeeVendingMachineTest();
        test.setUp(); test.selectsTheExistingDecoratedCoffee();
        test.setUp(); test.paidDispenseConsumesTheSelectedRecipe();
        test.setUp(); test.insufficientPaymentDoesNotConsumeIngredientsAndCanBeCancelled();
        test.setUp(); test.refillUsesTheExistingInventoryApi();
        System.out.println("PASS four native coffee API tests");
    }
}
