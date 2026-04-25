package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Expense struct {
	ID     int     `json:"id"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}

type Store struct {
	mu       sync.Mutex
	expenses []Expense
	nextId   int
}

func NewStore() *Store {
	return &Store{
		nextId: 1,
		expenses: []Expense{
			{1, "Lunch", 12.50},
			{2, "Grab", 8.00},
		},
	}
}

func (s *Store) GetAll() []Expense {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Expense{}, s.expenses...)
}

func (s *Store) Add(desc string, amount float64) Expense {
	s.mu.Lock()
	defer s.mu.Unlock()

	e := Expense{ID: s.nextId, Desc: desc, Amount: amount}
	s.expenses = append(s.expenses, e)
	s.nextId++
	return e
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, e := range s.expenses {
		if e.ID == id {
			fmt.Println(s.expenses[:i])
			fmt.Println(s.expenses[i+1:])

			s.expenses = append(s.expenses[:i], s.expenses[i+1:]...)
			return true
		}
	}

	return false
}

func main() {
	store := NewStore()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(store.GetAll())
		case http.MethodPost:
			var input struct {
				Desc   string
				Amount float64
			}

			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				log.Println(err)
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}

			e := store.Add(input.Desc, input.Amount)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(e)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/expenses/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var id int
		_, err := fmt.Sscanf(r.URL.Path, "/expenses/%d", &id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if store.Delete(id) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	})

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
