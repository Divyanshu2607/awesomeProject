package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("worker:", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		result <- job * 2 // Corrected processing: send `job * 2` instead of `id * 2`
	}
}

func main() {
	jobs := make(chan int, 5)
	result := make(chan int, 5) // Ensure it's large enough
	var wg sync.WaitGroup

	// Start 3 workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, result, &wg)
	}

	// Send jobs
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs) // Close job channel so workers stop reading

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(result) // Close result channel **after all workers are done**
	}()

	// Read results
	for res := range result {
		fmt.Println("Result:", res)
	}
}
