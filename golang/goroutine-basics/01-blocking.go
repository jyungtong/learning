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
	fmt.Println("main: before slowTask(1)")
	slowTask(1)
	fmt.Println("main: after slowTask(1)")
}
