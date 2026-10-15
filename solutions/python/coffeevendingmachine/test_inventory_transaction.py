import unittest
from concurrent.futures import ThreadPoolExecutor
from unittest.mock import patch

from coffee_vending_machine import CoffeeVendingMachine
from enums import CoffeeType, Ingredient
from inventory import Inventory
from states import OutOfIngredientState, ReadyState


class InventoryTransactionTest(unittest.TestCase):
    def setUp(self):
        self.inventory = Inventory.get_instance()
        self.machine = CoffeeVendingMachine.get_instance()
        with self.inventory._lock:
            self.inventory._stock.clear()
        self.machine.reset()
        self.machine.set_state(ReadyState())

    def tearDown(self):
        self.machine.reset()
        self.machine.set_state(ReadyState())

    def test_rejected_deduction_does_not_prepare_or_keep_payment(self):
        coffee_type = CoffeeType.ESPRESSO
        recipe = {
            Ingredient.COFFEE_BEANS: 7,
            Ingredient.WATER: 30,
        }
        for ingredient, amount in recipe.items():
            self.inventory.add_stock(ingredient, amount)
        self.machine.select_coffee(coffee_type, [])
        self.machine.insert_money(150)
        coffee = self.machine.get_selected_coffee()

        self.assertTrue(self.inventory.deduct_ingredients(recipe))
        with patch.object(coffee, "prepare") as prepare:
            self.machine.dispense_coffee()
        prepare.assert_not_called()
        self.assertEqual(self.machine.get_money_inserted(), 0)
        self.assertIsInstance(self.machine.get_state(), OutOfIngredientState)

    def test_only_one_concurrent_deduction_can_claim_last_recipe(self):
        recipe = {
            Ingredient.COFFEE_BEANS: 7,
            Ingredient.WATER: 30,
        }
        for ingredient, amount in recipe.items():
            self.inventory.add_stock(ingredient, amount)

        with ThreadPoolExecutor(max_workers=2) as pool:
            results = list(pool.map(lambda _: self.inventory.deduct_ingredients(recipe), range(2)))
        self.assertEqual(results.count(True), 1)
        self.assertEqual(results.count(False), 1)


if __name__ == "__main__":
    unittest.main()
