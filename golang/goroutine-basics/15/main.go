package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
