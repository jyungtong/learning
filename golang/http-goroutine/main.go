package main

import (
	"context"
	"expense-tracker/internal/api"
	"expense-tracker/internal/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	expenseStore := store.NewStore()
	defer expenseStore.Close()

	handler := api.NewHandler(expenseStore)

	server := &http.Server{
		Addr:    ":8080",
		Handler: api.LoggingMiddleware(handler),
	}

	go logExpenseCount(ctx, expenseStore)

	go func() {
		log.Println("listening :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}

func logExpenseCount(ctx context.Context, expenseStore *store.Store) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count, err := expenseStore.Count(ctx)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("expense count: %d", count)
		case <-ctx.Done():
			log.Printf("stopping expense count ticker")
			return
		}
	}
}
