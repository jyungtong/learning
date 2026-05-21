package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		requests  = 10
		burst     = 3
		rateEvery = 200 * time.Millisecond
	)

	requestsCh := make(chan int, requests)
	for i := 1; i < requests; i++ {
		requestsCh <- i
	}
	close(requestsCh)

	limiter := make(chan time.Time, burst)

	for i := 0; i < burst; i++ {
		limiter <- time.Now()
	}

	stopRefill := make(chan struct{})
	go func() {
		ticker := time.NewTicker(rateEvery)
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				select {
				case limiter <- t:
					fmt.Println("token added")
				default:
					fmt.Println("bucket full, drop refill")
				}
			case <-stopRefill:
				return
			}
		}
	}()

	fmt.Println("=== Section A: Backpressure (wait for token) ===")

	requestsA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, req := range requestsA {
		tokenTime := <-limiter
		fmt.Printf("request %d processed at %s (token from %s)\n", req, time.Now().Format("15:04:05.000"), tokenTime.Format("15:04:05.000"))
	}

	fmt.Println("\n=== Section B: Timeout guard (drop if no token in 150ms) ===")

	requestsB := []int{101, 102, 103, 104, 105, 106, 107, 108, 109, 110}
	for _, req := range requestsB {
		select {
		case <-limiter:
			fmt.Printf("request %d processed at %s\n", req, time.Now().Format("15:04:05.000"))
		case <-time.After(150 * time.Millisecond):
			fmt.Printf("request %d dropped\n", req)
		}
	}
	
	close(stopRefill)

	fmt.Println("\nDone")
}
