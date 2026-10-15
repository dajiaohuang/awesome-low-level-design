package parkinglot

import (
	"sync"
	"testing"
)

func TestConcurrentParkVehicleAllocatesSpotOnce(t *testing.T) {
	lot := &ParkingLot{}
	lot.AddLevel(NewLevel(1, 2))

	start := make(chan struct{})
	successes := make(chan Vehicle, 32)
	var workers sync.WaitGroup
	for index := 0; index < cap(successes); index++ {
		vehicle := NewMotorcycle("REVIEW-357-" + string(rune('A'+index)))
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			if lot.ParkVehicle(vehicle) {
				successes <- vehicle
			}
		}()
	}
	close(start)
	workers.Wait()
	close(successes)

	var winner Vehicle
	count := 0
	for vehicle := range successes {
		winner = vehicle
		count++
	}
	if count != 1 {
		t.Fatalf("successful concurrent parks = %d, want 1", count)
	}
	if !lot.UnparkVehicle(winner) {
		t.Fatal("failed to unpark the winning vehicle")
	}
	if !lot.ParkVehicle(NewMotorcycle("REVIEW-357-next")) {
		t.Fatal("spot was not available after unpark")
	}
}
