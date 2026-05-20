package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string) {
	msgs := []string{"a","b","c"}

	for _, m := range msgs {
		ch <- m
		time.Sleep(200 * time.Millisecond)
	}

	close(ch)
}

func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Println("received:", msg)
	}
}

func main() {
	ch := make(chan string)

	go producer(ch)
	consumer(ch)
}

