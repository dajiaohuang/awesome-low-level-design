package singleton

import "sync"

// DoubleCheckedSingleton implements double-checked locking singleton pattern
type DoubleCheckedSingleton struct{}

var (
	doubleCheckedInstance *DoubleCheckedSingleton
	dcOnce                sync.Once
)

// GetDoubleCheckedInstance returns the singleton instance
func GetDoubleCheckedInstance() *DoubleCheckedSingleton {
	dcOnce.Do(func() {
		doubleCheckedInstance = &DoubleCheckedSingleton{}
	})
	return doubleCheckedInstance
}
