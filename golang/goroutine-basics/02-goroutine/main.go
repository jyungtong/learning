package main

import (
	"fmt"
	"time"
)

func slowTask(id int) {
	for i:= 1; i <= 3; i++ {
		fmt.Printf("task %d: step %d\n", id, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main: starting goroutines")
	go slowTask(1)
	go slowTask(2)

	fmt.Println("main: goroutines launched, sleeping to let them run")
	time.Sleep(2 * time.Second)

	fmt.Println("main: done")
}
