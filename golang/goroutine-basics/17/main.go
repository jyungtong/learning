package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int)  {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main()  {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a < numJobs; a++ {
		fmt.Println("result:", <-results)
	}
}
