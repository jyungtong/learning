package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func slowDb(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
	defer cancel()

	err := slowDb(ctx)
	if err != nil {
		log.Printf("err: %v", err)
		http.Error(w, "request timed out", http.StatusGatewayTimeout)
		return
	}

	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
