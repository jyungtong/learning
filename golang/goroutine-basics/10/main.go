package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		fast <- "fast ready"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		fast <- "slow ready"
	}()

	select {
	case msg := <-fast:
		fmt.Println("first:", msg)
	case msg := <-slow:
		fmt.Println("first:", msg)
	}

	select {
	case msg := <-fast:
		fmt.Println("second:", msg)
	case msg := <-slow:
		fmt.Println("second:", msg)
	}
}

