package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

var (
	store = map[string]string{}
	mu sync.RWMutex
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody struct {
		URL string
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		log.Println("err decoding r.Body:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if reqBody.URL == "" {
		log.Println("url empty: ", reqBody.URL)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	code := generateShortCode()
	mu.Lock()
	store[code] = reqBody.URL
	mu.Unlock()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/%s", code)})
}

func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	mu.RLock()
	url := store[code]
	mu.RUnlock()
	if url == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Write([]byte(store[code]))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", shortenHandler)
	mux.HandleFunc("/{code}", getUrlHandler)

	handler := loggingMiddleware(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
