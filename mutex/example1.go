package mutex

import (
	"fmt"
	"sync"
	"time"
)

// SharedCounter demonstrates a shared counter protected by a mutex.
type SharedCounter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter.
func (c *SharedCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely retrieves the counter's value.
func (c *SharedCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func RunMutexConcurrencyExample() {
	counter := SharedCounter{}
	var wg sync.WaitGroup

	// Simulate concurrent access to the counter.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			time.Sleep(time.Millisecond) // Simulate some work.
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
