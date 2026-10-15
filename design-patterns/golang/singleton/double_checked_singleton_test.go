package singleton

import (
	"sync"
	"testing"
)

func TestDoubleCheckedSingletonConcurrentFirstAccess(t *testing.T) {
	const callers = 64
	instances := make([]*DoubleCheckedSingleton, callers)
	var wg sync.WaitGroup
	wg.Add(callers)
	for i := range instances {
		go func(i int) {
			defer wg.Done()
			instances[i] = GetDoubleCheckedInstance()
		}(i)
	}
	wg.Wait()
	for i := 1; i < len(instances); i++ {
		if instances[i] != instances[0] {
			t.Fatal("concurrent callers received distinct singleton instances")
		}
	}
}
