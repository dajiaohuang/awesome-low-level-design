package onlineshopping

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type OnlineShoppingService struct {
	Users    map[string]*User
	Products map[string]*Product
	Orders   map[string]*Order
	mu       sync.RWMutex
	nextOrderID uint64
}

var instance *OnlineShoppingService
var instanceOnce sync.Once

func GetInstance() *OnlineShoppingService {
	instanceOnce.Do(func() {
		instance = &OnlineShoppingService{
			Users:    make(map[string]*User),
			Products: make(map[string]*Product),
			Orders:   make(map[string]*Order),
		}
	})
	return instance
}

func (s *OnlineShoppingService) RegisterUser(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Users[user.ID] = user
}

func (s *OnlineShoppingService) AddProduct(product *Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Products[product.ID] = product
}

func (s *OnlineShoppingService) SearchProducts(keyword string) []*Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var results []*Product
	for _, product := range s.Products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
			results = append(results, product)
		}
	}
	return results
}

func (s *OnlineShoppingService) PlaceOrder(user *User, cart *ShoppingCart, paymentMethod Payment) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var orderItems []OrderItem
	items := cart.GetItems()
	for _, item := range items {
		if !item.Product.IsAvailable(item.Quantity) {
			return nil, errors.New("product unavailable")
		}
	}
	for _, item := range items {
		item.Product.UpdateQuantity(-item.Quantity)
		orderItems = append(orderItems, item)
	}

	if len(orderItems) == 0 {
		return nil, errors.New("no available products in the cart")
	}

	s.nextOrderID++
	order := NewOrder(fmt.Sprintf("ORDER-%s-%d", user.ID, s.nextOrderID), user, orderItems)
	s.Orders[order.ID] = order
	user.AddOrder(*order)
	cart.Clear()

	if paymentMethod.ProcessPayment(order.TotalAmount) {
		order.SetStatus(Processing)
	} else {
		order.SetStatus(Cancelled)
		for _, item := range orderItems {
			item.Product.UpdateQuantity(item.Quantity)
		}
	}

	return order, nil
}
