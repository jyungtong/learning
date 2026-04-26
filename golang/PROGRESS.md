# Go Backend Learning Progress

Learner:
- Experienced TS/Node/AWS backend dev.
- Learning Go backend/microservice style.
- Prefers build-first, explain pattern after.
- Caveman mode active: terse, exact.

## Project

Expense Tracker API.

## Syllabus

| Day | Topic | Status |
| --- | --- | --- |
| 1 | `net/http`, structs, JSON encode | Completed |
| 2 | POST body decode, pointer `&`, `append` | Completed |
| 3 | `sync.Mutex`, `defer`, pointer receivers, Delete | Completed |
| 4 | PostgreSQL via `pgx` | Completed |
| 5 | Middleware, structured errors | Completed |
| 6 | Project structure, packages | Completed |
| 7 | Goroutines, ticker, graceful shutdown | Completed |

Current Day 6 folder:
- `~/Documents/dev/personal/learning/golang/http-file-structure`

Current Day 7 folder:
- `~/Documents/dev/personal/learning/golang/http-goroutine`

Previous Day 5 folder:
- `~/Documents/dev/personal/learning/golang/http-pgsql`

Older Day 3 folder:
- `~/Documents/dev/personal/learning/golang/http-sync`

Important:
- `http-sync` is still old in-memory version.
- `http-pgsql` is Day 5 single-file version.
- `http-file-structure` is Day 6 split-package version with `internal/api` and `internal/store`.
- `http-goroutine` is Day 7 version with goroutines, ticker, and graceful shutdown.

## Completed

Day 1-3:
- Basic Go HTTP API.
- In-memory expense store.
- Routes:
  - `GET /health`
  - `GET /expenses`
  - `POST /expenses`
  - `DELETE /expenses/:id`

Day 4:
- Postgres added with Docker Compose.
- `pgxpool` store added.
- Store methods use `context.Context`:
  - `GetAll(ctx)`
  - `Add(ctx, desc, amount)`
  - `Delete(ctx, id)`
- Table auto-created on startup.

Day 5:
- Logging middleware added.
- Explicit `http.NewServeMux` added.
- `writeJSON` helper added.
- `writeError` helper added.
- Structured JSON errors added with `{"error":"..."}` shape.
- Error paths return immediately.
- Delete success returns JSON.
- `go run .` verified from `http-pgsql`.
- Curl checks verified.

Day 6:
- App split from one `main.go` into packages.
- `main.go` now only wires dependencies and starts server.
- `internal/store/store.go` owns Postgres access and `Expense` model.
- `internal/api/handlers.go` owns route registration and handlers.
- `internal/api/response.go` owns JSON/error response helpers.
- `internal/api/middleware.go` owns logging middleware.
- `api.NewHandler(store *store.Store) http.Handler` returns mux as interface.
- `api.LoggingMiddleware` exported for `main` package use.
- `writeJSON` and `writeError` stay unexported inside `api` package.
- `gofmt -w .` run from `http-file-structure`.
- `go test ./...` compile check passed from `http-file-structure`.

Day 7:
- `main.go` uses `http.Server` instead of `http.ListenAndServe`.
- Server starts in goroutine.
- Background ticker goroutine logs expense count every 10 seconds.
- `signal.NotifyContext` catches Ctrl+C and SIGTERM.
- `<-ctx.Done()` waits for shutdown signal.
- `server.Shutdown` with 5-second timeout context.
- `store.Count(ctx)` added for ticker DB work.
- Ticker exits cleanly via `ctx.Done()` in `select`.
- `http.ErrServerClosed` ignored on shutdown.
- `gofmt -w .` run from `http-goroutine`.
- `go test ./...` compile check passed from `http-goroutine`.

## Current Code State

Files:
- `~/Documents/dev/personal/learning/golang/http-goroutine/main.go`
- `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/handlers.go`
- `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/middleware.go`
- `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/response.go`
- `~/Documents/dev/personal/learning/golang/http-goroutine/internal/store/store.go`

Current state after Day 7:
- Same API behavior as Day 6.
- `main.go` starts server in goroutine with `http.Server`.
- `main.go` starts ticker goroutine via `go logExpenseCount(ctx, expenseStore)`.
- `main.go` waits for shutdown signal via `<-ctx.Done()`.
- `main.go` calls `server.Shutdown(shutdownCtx)` with 5-second timeout.
- `logExpenseCount` uses `time.NewTicker` and `select` with `ctx.Done()`.
- `store.Count(ctx)` returns expense count for ticker.
- `go test ./...` passes from `http-goroutine`.

## Day 7 Goal

Add lifecycle and concurrency patterns:
- goroutines for server and background work
- ticker for periodic tasks
- graceful shutdown via OS signals
- context cancellation for goroutine coordination
- `http.Server` for controlled shutdown

Keep project structure same as Day 6.

## Day 5 Goal

Add production-ish HTTP patterns:
- logging middleware
- explicit mux
- JSON response helper
- structured JSON errors
- central error response helper

Keep project single-file until Day 6 project structure.

## Day 5 Target Shape

Add:

```go
type ErrorResponse struct {
	Error string `json:"error"`
}
```

Add helpers:

```go
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
```

Add middleware:

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
```

Use explicit mux:

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", ...)
mux.HandleFunc("/expenses", ...)
mux.HandleFunc("/expenses/", ...)

log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
```

## Day 5 Checklist

- All API responses JSON except `/health` can stay plain `OK`.
- All errors use `{"error":"..."}` shape.
- No `http.Error` left in handlers.
- Every error response returns immediately.
- Middleware logs each request.
- Server uses explicit mux.
- `go run .` works from `http-pgsql`.
- Curl checks pass:

```bash
curl -i http://localhost:8080/expenses
curl -i -X POST http://localhost:8080/expenses -d 'bad'
curl -i -X DELETE http://localhost:8080/expenses/not-a-number
```

## Day 7 Teaching Points

- `go fn()` starts new goroutine.
- Goroutines are lightweight threads managed by Go runtime.
- `time.NewTicker(interval)` fires events at regular intervals.
- `defer ticker.Stop()` prevents resource leak.
- `select` waits on multiple channels, picks first ready.
- `context.Context` carries cancellation across goroutines.
- `signal.NotifyContext` converts OS signals to context cancellation.
- `http.Server.Shutdown` stops accepting new requests, drains active ones.
- `http.ErrServerClosed` is expected on graceful shutdown, not an error.
- Shutdown context must be fresh, not already-canceled signal context.
- `<-ctx.Done()` blocks until context is canceled.
- `context.WithTimeout` creates deadline for shutdown operation.

## Teaching Points

- `any` is alias for `interface{}`.
- One JSON response path reduces duplicated headers/encoder code.
- Structured errors keep clients predictable.
- Headers must be set before `WriteHeader`.
- Middleware wraps `http.Handler`.
- `http.Handler` interface has `ServeHTTP`.
- `http.HandlerFunc` adapts a function to handler.
- Request flow: server -> middleware -> mux -> handler.
- Go package usually maps to folder.
- Import path starts with module name from `go.mod`: `expense-tracker`.
- `internal/...` packages are importable only from inside the same module tree.
- Capitalized identifiers are exported across packages.
- Lowercase identifiers are private to package.
- Interfaces are satisfied implicitly; no `implements` keyword.
- `http.Handler` requires `ServeHTTP(ResponseWriter, *Request)`.
- `*http.ServeMux` satisfies `http.Handler`, so `return mux` works from `func NewHandler(...) http.Handler`.
- Returning `http.Handler` hides concrete mux implementation from caller.

## Next Step

Start Day 8: (next topic TBD)
