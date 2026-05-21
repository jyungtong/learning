package main

import (
	"fmt"
	"runtime"
	"time"
)

func leaky() {
	ch := make(chan int)
	go func ()  {
		ch <- 1
		fmt.Println("leaky goroutine: sent (you'll never see this)")
	}()
}

func main() {
	fmt.Println("goroutines before:", runtime.NumGoroutine())

	leaky()

	// give scheduler time to start goroutines
	time.Sleep(100 * time.Millisecond)

	fmt.Println("goroutines after leaky():", runtime.NumGoroutine())
	fmt.Println("main: done (leaked goroutine still stuck)")
}
