package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second"

	// ch <- "third"

	go func() {
		time.Sleep(500 * time.Millisecond)

		ch <- "third"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// time.Sleep(100 * time.Millisecond)
	fmt.Println(<-ch)
}
