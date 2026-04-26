package api

import (
	"encoding/json"
	"expense-tracker/internal/store"
	"fmt"
	"log"
	"net/http"
)

func NewHandler(store *store.Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"msg": "OK"})
	})

	mux.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			expenses, err := store.GetAll(r.Context())
			if err != nil {
				log.Println(err)
				writeError(w, http.StatusInternalServerError, "internal server error")
				return
			}

			writeJSON(w, http.StatusOK, expenses)
		case http.MethodPost:
			var input struct {
				Desc   string
				Amount float64
			}
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				writeError(w, http.StatusBadRequest, "bad request")
				return
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
			writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
		} else {
			writeError(w, http.StatusNotFound, "not found")
		}
	})

	return mux //TODO: ask why diff type works
}
