package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Expense struct {
	ID     int     `json:"id"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}

type Store struct {
	db *pgxpool.Pool
}

func NewStore() *Store {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://expense_user:expense_password@localhost:5432/expense_tracker"
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			description TEXT NOT NULL,
			amount DOUBLE PRECISION NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{db}
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) GetAll(ctx context.Context) ([]Expense, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, description, amount
		FROM expenses
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.Desc, &e.Amount); err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *Store) Add(ctx context.Context, desc string, amount float64) (Expense, error) {
	var e Expense

	err := s.db.QueryRow(ctx, `
		INSERT INTO expenses (description, amount)
		VALUES ($1, $2)
		RETURNING id, description, amount
	`, desc, amount).Scan(&e.ID, &e.Desc, &e.Amount)

	if err != nil {
		return Expense{}, err
	}

	return e, nil
}

func (s *Store) Delete(ctx context.Context, id int) (bool, error) {
	result, err := s.db.Exec(ctx, `
		DELETE FROM expenses
		WHERE id = $1
	`, id)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func main() {
	store := NewStore()
	defer store.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			expenses, err := store.GetAll(r.Context())
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expenses)
		case http.MethodPost:
			var input struct{
				Desc string
				Amount float64
			}
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
			}

			e, err := store.Add(r.Context(), input.Desc, input.Amount)
			if err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

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

		deleted, err := store.Delete(r.Context(), id)

		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if deleted {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	})

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
