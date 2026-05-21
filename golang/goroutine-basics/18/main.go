package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(2 * time.Second) // slow worker
		results <- j * 2
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 2

	jobs := make(chan int, 2) //small buffer force timeout
	results := make(chan int, 2)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	sent := 0
	dropped := 0

	// send jobs without timeout
	for j := 1; j < numJobs; j++ {
		select {
		case jobs <- j:
			sent++
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("timeout sending job %d\n", j)
			dropped++
		}
	}
	close(jobs)

	// background: close results when workers Done
	go func() {
		wg.Wait()
		close(results)
	}()

	received := 0

	// collect results with timeout
	for {
		select {
		case r, ok := <-results:
			if !ok {
				fmt.Println("results chnnel closed")
				goto done
			}
			fmt.Println("result:", r)
			received++
		case <-time.After(5 * time.Second):
			fmt.Println("timeout waiting for results")
			goto done
		}
	}

done:
	fmt.Printf("sent: %d, dropped: %d, received: %d\n", sent, dropped, received)
}
