package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("receiver: waitign...")
		msg := <-ch
		fmt.Println("receiver: got", msg)
	}()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main: about to send")
	ch <- "hello"
	fmt.Println("main: send completed")

	time.Sleep(time.Second)
}
