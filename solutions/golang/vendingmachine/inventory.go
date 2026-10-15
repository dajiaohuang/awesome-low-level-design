package vending_machine

import "sync"

type Inventory struct {
	products map[*Product]int
	mu       sync.RWMutex
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (inv *Inventory) AddProduct(product *Product, quantity int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.products[product] = quantity
}

func (inv *Inventory) IsAvailable(product *Product) bool {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	qty, exists := inv.products[product]
	return exists && qty > 0
}

func (inv *Inventory) Dispense(product *Product) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.products[product]--
}
