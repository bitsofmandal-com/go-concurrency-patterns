package workerpool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func RunWorkerPoolConcurrencyExample() {
	numTasks := 20
	numWorkers := 5
	taskChan := make(chan Task, numTasks)
	var wg sync.WaitGroup
	// Start the workers.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskChan, &wg)

	}
	// Send tasks to the workers.
	for i := 1; i <= numTasks; i++ {
		taskChan <- Task{ID: i}
	}
	// Close the tasks channel to signal that no more tasks will be sent.
	close(taskChan)
	// Wait for all workers to finish.
	wg.Wait()
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		// Simulate some work.
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
	}
}
