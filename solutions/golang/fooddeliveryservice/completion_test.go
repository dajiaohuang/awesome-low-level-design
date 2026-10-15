package fooddeliveryservice

import (
 "testing"
 "time"
)

func TestCompletedAgentReuse(t *testing.T) {
	s := GetFoodDeliveryService()
	c := NewCustomer("C001", "Customer", "c@example.com", "123")
	m := NewMenuItem("M001", "Meal", "Meal", 10)
	r := NewRestaurant("R001", "Restaurant", "Address", []*MenuItem{m})
	d := NewDeliveryAgent("D001", "Agent", "456")
	s.RegisterCustomer(c)
	s.RegisterRestaurant(r)
	s.RegisterDeliveryAgent(d)
	ready := func() *Order {
        time.Sleep(2 * time.Millisecond)
		o, err := s.PlaceOrder(c.ID, r.ID, []*OrderItem{NewOrderItem(m, 1)})
		if err != nil {
			t.Fatal(err)
		}
		if err = s.UpdateOrderStatus(o.ID, OrderStatusConfirmed); err != nil {
			t.Fatal(err)
		}
		return o
	}
	first := ready()
	if first.DeliveryAgent != d {
		t.Fatal("first assignment")
	}
	s.UpdateOrderStatus(first.ID, OrderStatusDelivered)
	if !d.IsAvailable() {
		t.Fatal("completion did not release")
	}
	second := ready()
	if second.DeliveryAgent != d {
		t.Fatal("agent not reused")
	}
	s.UpdateOrderStatus(first.ID, OrderStatusDelivered)
	if d.IsAvailable() {
		t.Fatal("stale completion released new assignment")
	}
	third := ready()
	if third.DeliveryAgent != nil {
		t.Fatal("busy agent reused")
	}
	s.UpdateOrderStatus(second.ID, OrderStatusDelivered)
	s.UpdateOrderStatus(third.ID, OrderStatusConfirmed)
	if third.DeliveryAgent != d {
		t.Fatal("second completion did not release")
	}
}
