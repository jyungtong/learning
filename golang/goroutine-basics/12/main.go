package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func slowWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("worker: timed out:", ctx.Err())
			return
		default:
			fmt.Println("worker: step", i)
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("worker: finsihed all steps")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go slowWorker(ctx, &wg)

	wg.Wait()
	fmt.Println("main: done")
}
