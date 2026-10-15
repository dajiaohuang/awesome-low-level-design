package vending_machine

import "testing"

func TestDispenseProductRemovesLastUnitFromAvailability(t *testing.T) {
	machine := NewVendingMachine()
	product := NewProduct("coffee", 0.25)
	machine.inventory.AddProduct(product, 1)

	machine.SelectProduct(product)
	machine.InsertCoin(QUARTER)
	machine.DispenseProduct()
	machine.ReturnChange()

	if machine.inventory.IsAvailable(product) {
		t.Fatal("sold-out product remains available")
	}
}
