package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

func hashFile(wg *sync.WaitGroup, f string, res chan<- string) {
	defer wg.Done()

	hasher := sha256.New()

	byteString := []byte(f)
	hasher.Write(byteString)
	res <- fmt.Sprintf("%x", hasher.Sum(nil))
	// fmt.Printf("%s: %x\n", f, hasher.Sum(nil))
}

func main() {
	res := make(chan string)
	var wg sync.WaitGroup

	files := []string{
		"/some/file1",
		"/some/file2",
		"/some/file3",
	}

	for _, f := range files {
		wg.Add(1)
		go hashFile(&wg, f, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Println(r)
	}

	fmt.Println("main: finished")
}
