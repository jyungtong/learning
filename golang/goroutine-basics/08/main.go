package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		msgs := []string{"a", "b", "c"}
		for _, m := range msgs {
			ch <- m
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("received:", msg)
	}

	fmt.Println("channel closed, range ended")
}

