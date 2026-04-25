package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Expense struct {
	ID     int     `json:"id"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Store struct {
	db *pgxpool.Pool
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println(err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
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

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{ "msg": "OK" })
	})

	mux.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			expenses, err := store.GetAll(r.Context())
			if err != nil {
				log.Println(err)
				writeError(w, http.StatusInternalServerError, "internal server error")
			}

			writeJSON(w, http.StatusOK, expenses)
		case http.MethodPost:
			var input struct{
				Desc string
				Amount float64
			}
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				writeError(w, http.StatusBadRequest, "bad request")
			}

			e, err := store.Add(r.Context(), input.Desc, input.Amount)
			if err != nil {
				log.Println(err)
				writeError(w, http.StatusInternalServerError, "internal server error")
				return
			}

			writeJSON(w, http.StatusCreated, e)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	mux.HandleFunc("/expenses/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		var id int
		_, err := fmt.Sscanf(r.URL.Path, "/expenses/%d", &id)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}

		deleted, err := store.Delete(r.Context(), id)

		if err != nil {
			log.Println(err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		if deleted {
			w.WriteHeader(http.StatusNoContent)
		} else {
			writeError(w, http.StatusNotFound, "not found")
		}
	})

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
}
