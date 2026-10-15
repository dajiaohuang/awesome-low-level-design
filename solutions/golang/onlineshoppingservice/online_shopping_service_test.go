package onlineshopping

import (
	"sync"
	"testing"
)

type acceptedPayment struct{}

func (acceptedPayment) ProcessPayment(float64) bool { return true }

func TestConcurrentOrdersCannotSellSameStock(t *testing.T) {
	product := NewProduct("one", "Limited item", "", 10, 1)
	service := &OnlineShoppingService{
		Users:    make(map[string]*User),
		Products: map[string]*Product{product.ID: product},
		Orders:   make(map[string]*Order),
	}
	users := []*User{
		NewUser("buyer-1", "Buyer 1", "one@example.com", ""),
		NewUser("buyer-2", "Buyer 2", "two@example.com", ""),
	}
	carts := []*ShoppingCart{NewShoppingCart(), NewShoppingCart()}
	for _, cart := range carts {
		cart.AddItem(product, 1)
	}

	var wg sync.WaitGroup
	results := make(chan error, len(users))
	for i, user := range users {
		wg.Add(1)
		go func(user *User, cart *ShoppingCart) {
			defer wg.Done()
			_, err := service.PlaceOrder(user, cart, acceptedPayment{})
			results <- err
		}(user, carts[i])
	}
	wg.Wait()
	close(results)

	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("successful orders = %d, want exactly 1", successes)
	}
	if product.Quantity != 0 {
		t.Fatalf("remaining stock = %d, want 0", product.Quantity)
	}
	if len(service.Orders) != 1 {
		t.Fatalf("recorded orders = %d, want 1", len(service.Orders))
	}
}
