package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Expense struct {
	ID     int     `json:"id"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}

var expenses = []Expense{
	{1, "Lunch", 12.50},
	{2, "Grab", 8.00},
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetExpenses(w)
		case http.MethodPost:
			handlePostExpenses(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGetExpenses(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func handlePostExpenses(w http.ResponseWriter, r *http.Request) {
	var input Expense
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	input.ID = len(expenses) + 1
	expenses = append(expenses, input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}
