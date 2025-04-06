package faninout

import (
	"fmt"
	"time"
)

func RunFanInFanOutConcurrencyExample() {
  jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Fan out: Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Fan in: Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}


func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
      fmt.Printf("Worker %d processing job %d\n", id, j)
      time.Sleep(time.Second) // Simulate work
      results <- j * 2        // Send result back
  }
}
