package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("worker: stopping")
			return
		case t := <-ticker.C:
			fmt.Println("worker: tick at", t.Format("15:04:05.000"))
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	time.Sleep(1 * time.Second)
	fmt.Println("main: stopping ticker")
	cancel()

	wg.Wait()
	fmt.Println("main: done")
}

