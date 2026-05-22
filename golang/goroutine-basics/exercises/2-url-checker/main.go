package main

import (
	"fmt"
	"net/http"
	"sync"
)

func urlCheck(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: error {%v}\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%s: %d\n", url, resp.StatusCode)
}

func main() {
	// res := make(chan string)
	var wg sync.WaitGroup

	urls := []string{
		"https://jsonplaceholder.typicode.com/todos",
		"https://jsonplaceholder.typicode.com/users",
		"https://sklfjlsjfslkfjsdfjsdflk.com",
		"https://jsonplaceholder.typicode.com/posts",
	}

	for _, url := range urls {
		wg.Add(1)
		go urlCheck(&wg, url)
	}

	// go func() {
		// close(res)
	// }()

	// for r := range res {
	// 	fmt.Println(r)
	// }

	wg.Wait()
	fmt.Println("main: finished")
}
