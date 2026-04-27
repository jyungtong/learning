package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 1; i <= 3; i++ {
		fmt.Println("hello", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go sayHello()

	fmt.Println("main keeps running")
	time.Sleep(4 * time.Second)
	fmt.Println("main done")
}
