package main

import (
	"fmt"
	"sync"
)

func fanOut(src <-chan int, n int) []<-chan int {
	channels := make([]<-chan int, n)
	for i := range n {
		ch := make(chan int)
		channels[i] = ch
		go func() {
			defer close(ch)
			for v := range src {
				ch <- v * v
			}
		}()
	}
	return channels
}

func fanIn(channels []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	jobs := make(chan int, 10)
	for i := 1; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	workers := fanOut(jobs, 3)
	results := fanIn(workers)

	for v := range results {
		fmt.Println(v)
	}
}
