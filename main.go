package main

import (
	"fmt"

	"github.com/bitsofmandal-com/workerpool"
	// "github.com/bitsofmandal-com/workerpool"
)

func main() {
	workerpool.RunWorkerPoolConcurrencyExample()
	fmt.Println("All tasks completed.")
}
