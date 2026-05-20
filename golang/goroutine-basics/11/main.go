package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("worker: context cancelled, exiting")
			return
		default:
			fmt.Println("worker: doing work")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: cancelling context")
	cancel()

	wg.Wait()
	fmt.Println("main: worker stopped cleanly")
}

