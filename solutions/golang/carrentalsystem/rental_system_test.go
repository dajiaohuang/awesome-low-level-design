package carrentalsystem

import (
	"testing"
	"time"
)

func TestSearchCarsIncludesDisjointReservationDates(t *testing.T) {
	system := &RentalSystem{
		cars:         make(map[string]*Car),
		reservations: make(map[string]*Reservation),
	}
	car := NewCar("Acme", "Roadster", 2026, "REVIEW-372", 100)
	system.AddCar(car)
	customer := &Customer{}
	firstStart := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	firstEnd := time.Date(2030, 1, 3, 0, 0, 0, 0, time.UTC)
	if _, err := system.MakeReservation(customer, car, firstStart, firstEnd); err != nil {
		t.Fatalf("first reservation: %v", err)
	}

	disjointStart := time.Date(2030, 1, 10, 0, 0, 0, 0, time.UTC)
	disjointEnd := time.Date(2030, 1, 12, 0, 0, 0, 0, time.UTC)
	got := system.SearchCars("Acme", "Roadster", disjointStart, disjointEnd)
	if len(got) != 1 || got[0] != car {
		t.Fatalf("search result = %#v, want the car for disjoint dates", got)
	}
}
