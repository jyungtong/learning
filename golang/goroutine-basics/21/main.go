package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func producer(ctx context.Context, jobs chan<- int, wg *sync.WaitGroup)  {
	defer wg.Done()
	defer close(jobs) // sender owns close

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	jobID := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("producer: shutdown signal, stop intake")
			return
		case <-ticker.C:
			select {
			case jobs <- jobID:
				fmt.Printf("producer: queued job %d\n", jobID)
				jobID++
			case <-ctx.Done():
				fmt.Println("producer: shutdown while queing")
				return
			}
		}
	}
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup)  {
	defer wg.Done()

	fmt.Printf("worker-%d: started\n", id)
	for job := range jobs {
		fmt.Printf("worker-%d: processing job %d\n", id, job)
		time.Sleep(700 * time.Millisecond) // simulate work
		fmt.Printf("worker-%d: finished job %d\n", id, job)
	}
	fmt.Printf("worker-%d: drained and exiting\n", id)
}

func main()  {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int, 8)
	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(1)
	go producer(ctx, jobs, &wg)

	for i := 1; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	fmt.Println("main: running (presss ctrl+c to shutdown)")
	<-ctx.Done()
	fmt.Println("main: signal received, waiting for graceful shutdown")

	wg.Wait()
	fmt.Println("main: shutdown complete")
}
