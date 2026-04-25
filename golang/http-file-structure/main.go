package main

import (
	"expense-tracker/internal/api"
	"expense-tracker/internal/store"
	"log"
	"net/http"
)

func main() {
	store := store.NewStore()
	defer store.Close()

	handler := api.NewHandler(store)

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", api.LoggingMiddleware(handler)))
}
