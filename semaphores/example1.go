package semaphores

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Semaphore controls concurrent access to a resource.
type Semaphore chan struct{}

// NewSemaphore creates a new semaphore with the given capacity.
func NewSemaphore(capacity int) Semaphore {
	return make(Semaphore, capacity)
}

// Acquire acquires a semaphore.
func (s Semaphore) Acquire() {
	s <- struct{}{}
}

// Release releases a semaphore.
func (s Semaphore) Release() {
	<-s
}

func RunSemaphoreConcurrencyExample() {
	numWorkers := 10
	semaphoreCapacity := 3 // Limit concurrent access to 3.

	semaphore := NewSemaphore(semaphoreCapacity)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, semaphore, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
}

// Worker simulates a task that requires access to a limited resource.
func worker(id int, semaphore Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()

	semaphore.Acquire()       // Acquire the semaphore before accessing the resource.
	defer semaphore.Release() // Release the semaphore after accessing the resource.

	fmt.Printf("Worker %d acquired semaphore, processing...\n", id)
	// Simulate resource-intensive work.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Worker %d released semaphore, finished.\n", id)

}
