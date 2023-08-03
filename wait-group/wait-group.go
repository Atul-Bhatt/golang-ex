package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Println("Job started", id)
	time.Sleep(time.Second)
	fmt.Println("Job finished", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers have finished")
}
