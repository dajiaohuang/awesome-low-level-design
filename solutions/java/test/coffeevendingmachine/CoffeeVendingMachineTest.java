package coffeevendingmachine;

import coffeevendingmachine.enums.CoffeeType;
import coffeevendingmachine.enums.Ingredient;
import coffeevendingmachine.enums.ToppingType;
import coffeevendingmachine.state.ReadyState;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.*;

public class CoffeeVendingMachineTest {
    private CoffeeVendingMachine machine;
    private Inventory inventory;

    @BeforeEach
    void setUp() {
        machine = CoffeeVendingMachine.getInstance();
        machine.cancel();
        inventory = Inventory.getInstance();
        for (Ingredient ingredient : Ingredient.values()) {
            inventory.addStock(ingredient, 500);
        }
    }

    @Test
    void selectsTheExistingDecoratedCoffee() {
        machine.selectCoffee(CoffeeType.LATTE, List.of(ToppingType.EXTRA_SUGAR));
        assertEquals("Latte, Extra Sugar", machine.getSelectedCoffee().getCoffeeType());
        assertEquals(230, machine.getSelectedCoffee().getPrice());
    }

    @Test
    void paidDispenseConsumesTheSelectedRecipe() {
        machine.selectCoffee(CoffeeType.LATTE, List.of());
        Map<Ingredient, Integer> recipe = machine.getSelectedCoffee().getRecipe();
        Map<Ingredient, Integer> before = machine.showIngredientsMap();
        machine.insertMoney(machine.getSelectedCoffee().getPrice());
        machine.dispenseCoffee();
        Map<Ingredient, Integer> after = machine.showIngredientsMap();
        recipe.forEach((ingredient, quantity) ->
            assertEquals(before.get(ingredient) - quantity, after.get(ingredient).intValue()));
        assertNull(machine.getSelectedCoffee());
        assertEquals(0, machine.getMoneyInserted());
        assertInstanceOf(ReadyState.class, machine.getState());
    }

    @Test
    void insufficientPaymentDoesNotConsumeIngredientsAndCanBeCancelled() {
        machine.selectCoffee(CoffeeType.ESPRESSO, List.of());
        Map<Ingredient, Integer> before = machine.showIngredientsMap();
        machine.insertMoney(1);
        machine.dispenseCoffee();
        assertEquals(before, machine.showIngredientsMap());
        assertEquals(1, machine.getMoneyInserted());
        machine.cancel();
        assertNull(machine.getSelectedCoffee());
        assertEquals(0, machine.getMoneyInserted());
        assertInstanceOf(ReadyState.class, machine.getState());
    }

    @Test
    void refillUsesTheExistingInventoryApi() {
        int before = machine.showIngredientsMap().get(Ingredient.WATER);
        inventory.addStock(Ingredient.WATER, 50);
        assertEquals(before + 50, machine.showIngredientsMap().get(Ingredient.WATER).intValue());
    }
}
