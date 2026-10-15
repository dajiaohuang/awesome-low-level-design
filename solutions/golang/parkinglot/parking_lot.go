package parkinglot

import "sync"

type ParkingLot struct {
	mu     sync.RWMutex
	levels []*Level
}

var (
	instance     *ParkingLot
	instanceOnce sync.Once
)

func GetParkingLotInstance() *ParkingLot {
	instanceOnce.Do(func() {
		instance = &ParkingLot{levels: []*Level{}}
	})
	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, level := range p.levels {
		if level.UnparkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) DisplayAvailability() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, level := range p.levels {
		level.DisplayAvailability()
	}
}
