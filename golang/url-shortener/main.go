package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

var (
	store = map[string]string{}
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
	// todo: generateShortCode and store 

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": generateShortCode()})
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
