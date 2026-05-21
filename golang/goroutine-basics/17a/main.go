package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main()  {
	var wg sync.WaitGroup

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("results:", r)
	}
}
