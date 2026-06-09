package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now().Format(time.DateOnly)
	fmt.Println(date)
}
