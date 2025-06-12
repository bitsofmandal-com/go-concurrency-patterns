package main

import (
	"fmt"
	"sync"
)

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			return
		}
		if num > 100 {
			close(ch)
			return
		}
		if (num % 2) == 0 {
			fmt.Printf("(even):\t%d\n", num)
		} else {
			fmt.Printf("(odd):\t%d\n", num)
		}
		ch <- num + 1
	}
}

func PrintSynchronizedCounter() {
	var wg sync.WaitGroup
	counter := 1
	ch := make(chan int, 1)

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(ch, &wg)
	}

	ch <- counter // Start the sequence
	wg.Wait()
}
