# Go test framework basics

**Session ID:** ses_14b417db9ffeG2C57SVx6aomdD
**Created:** 6/11/2026, 11:34:04 AM
**Updated:** 6/11/2026, 11:48:38 AM

---

## User

i want to learn go test framework, start from basic

---

## Assistant (Plan · Claude Opus 4.8 · 4.8s)

_Thinking:_

 The user is looking to get started with Go's testing framework, and I'm in planning mode so I'm just observing. Let me see what's currently in the workspace.



**Tool: bash**

**Input:**
```json
{
  "command": "rtk ls -la /home/jyung/Documents/dev/personal/learning/golang/ && echo \"---\" && rtk read /home/jyung/Documents/dev/personal/learning/golang/go.mod 2>/dev/null; go version 2>/dev/null",
  "description": "Inspect golang dir and go version"
}
```

**Output:**
```
1-creating-a-module/
2-multi-module-workspace/
date/
file/
goroutine-basics/
http-basic/
http-file-structure/
http-goroutine/
http-pgsql/
http-sync/
sqlite/
url-shortener/
PROGRESS.md  7.8K
basic.go  1.1K
---
go version go1.26.4 linux/amd64

```

---

## Assistant (Plan · Claude Opus 4.8 · 6.4s)

**Tool: read**

**Input:**
```json
{
  "filePath": "/home/jyung/Documents/dev/personal/learning/golang/PROGRESS.md"
}
```

**Output:**
```
<path>/home/jyung/Documents/dev/personal/learning/golang/PROGRESS.md</path>
<type>file</type>
<content>
1: # Go Backend Learning Progress
2: 
3: Learner:
4: - Experienced TS/Node/AWS backend dev.
5: - Learning Go backend/microservice style.
6: - Prefers build-first, explain pattern after.
7: - Caveman mode active: terse, exact.
8: 
9: ## Project
10: 
11: Expense Tracker API.
12: 
13: ## Syllabus
14: 
15: | Day | Topic | Status |
16: | --- | --- | --- |
17: | 1 | `net/http`, structs, JSON encode | Completed |
18: | 2 | POST body decode, pointer `&`, `append` | Completed |
19: | 3 | `sync.Mutex`, `defer`, pointer receivers, Delete | Completed |
20: | 4 | PostgreSQL via `pgx` | Completed |
21: | 5 | Middleware, structured errors | Completed |
22: | 6 | Project structure, packages | Completed |
23: | 7 | Goroutines, ticker, graceful shutdown | Completed |
24: 
25: Current Day 6 folder:
26: - `~/Documents/dev/personal/learning/golang/http-file-structure`
27: 
28: Current Day 7 folder:
29: - `~/Documents/dev/personal/learning/golang/http-goroutine`
30: 
31: Previous Day 5 folder:
32: - `~/Documents/dev/personal/learning/golang/http-pgsql`
33: 
34: Older Day 3 folder:
35: - `~/Documents/dev/personal/learning/golang/http-sync`
36: 
37: Important:
38: - `http-sync` is still old in-memory version.
39: - `http-pgsql` is Day 5 single-file version.
40: - `http-file-structure` is Day 6 split-package version with `internal/api` and `internal/store`.
41: - `http-goroutine` is Day 7 version with goroutines, ticker, and graceful shutdown.
42: 
43: ## Completed
44: 
45: Day 1-3:
46: - Basic Go HTTP API.
47: - In-memory expense store.
48: - Routes:
49:   - `GET /health`
50:   - `GET /expenses`
51:   - `POST /expenses`
52:   - `DELETE /expenses/:id`
53: 
54: Day 4:
55: - Postgres added with Docker Compose.
56: - `pgxpool` store added.
57: - Store methods use `context.Context`:
58:   - `GetAll(ctx)`
59:   - `Add(ctx, desc, amount)`
60:   - `Delete(ctx, id)`
61: - Table auto-created on startup.
62: 
63: Day 5:
64: - Logging middleware added.
65: - Explicit `http.NewServeMux` added.
66: - `writeJSON` helper added.
67: - `writeError` helper added.
68: - Structured JSON errors added with `{"error":"..."}` shape.
69: - Error paths return immediately.
70: - Delete success returns JSON.
71: - `go run .` verified from `http-pgsql`.
72: - Curl checks verified.
73: 
74: Day 6:
75: - App split from one `main.go` into packages.
76: - `main.go` now only wires dependencies and starts server.
77: - `internal/store/store.go` owns Postgres access and `Expense` model.
78: - `internal/api/handlers.go` owns route registration and handlers.
79: - `internal/api/response.go` owns JSON/error response helpers.
80: - `internal/api/middleware.go` owns logging middleware.
81: - `api.NewHandler(store *store.Store) http.Handler` returns mux as interface.
82: - `api.LoggingMiddleware` exported for `main` package use.
83: - `writeJSON` and `writeError` stay unexported inside `api` package.
84: - `gofmt -w .` run from `http-file-structure`.
85: - `go test ./...` compile check passed from `http-file-structure`.
86: 
87: Day 7:
88: - `main.go` uses `http.Server` instead of `http.ListenAndServe`.
89: - Server starts in goroutine.
90: - Background ticker goroutine logs expense count every 10 seconds.
91: - `signal.NotifyContext` catches Ctrl+C and SIGTERM.
92: - `<-ctx.Done()` waits for shutdown signal.
93: - `server.Shutdown` with 5-second timeout context.
94: - `store.Count(ctx)` added for ticker DB work.
95: - Ticker exits cleanly via `ctx.Done()` in `select`.
96: - `http.ErrServerClosed` ignored on shutdown.
97: - `gofmt -w .` run from `http-goroutine`.
98: - `go test ./...` compile check passed from `http-goroutine`.
99: 
100: ## Current Code State
101: 
102: Files:
103: - `~/Documents/dev/personal/learning/golang/http-goroutine/main.go`
104: - `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/handlers.go`
105: - `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/middleware.go`
106: - `~/Documents/dev/personal/learning/golang/http-goroutine/internal/api/response.go`
107: - `~/Documents/dev/personal/learning/golang/http-goroutine/internal/store/store.go`
108: 
109: Current state after Day 7:
110: - Same API behavior as Day 6.
111: - `main.go` starts server in goroutine with `http.Server`.
112: - `main.go` starts ticker goroutine via `go logExpenseCount(ctx, expenseStore)`.
113: - `main.go` waits for shutdown signal via `<-ctx.Done()`.
114: - `main.go` calls `server.Shutdown(shutdownCtx)` with 5-second timeout.
115: - `logExpenseCount` uses `time.NewTicker` and `select` with `ctx.Done()`.
116: - `store.Count(ctx)` returns expense count for ticker.
117: - `go test ./...` passes from `http-goroutine`.
118: 
119: ## Day 7 Goal
120: 
121: Add lifecycle and concurrency patterns:
122: - goroutines for server and background work
123: - ticker for periodic tasks
124: - graceful shutdown via OS signals
125: - context cancellation for goroutine coordination
126: - `http.Server` for controlled shutdown
127: 
128: Keep project structure same as Day 6.
129: 
130: ## Day 5 Goal
131: 
132: Add production-ish HTTP patterns:
133: - logging middleware
134: - explicit mux
135: - JSON response helper
136: - structured JSON errors
137: - central error response helper
138: 
139: Keep project single-file until Day 6 project structure.
140: 
141: ## Day 5 Target Shape
142: 
143: Add:
144: 
145: ```go
146: type ErrorResponse struct {
147: 	Error string `json:"error"`
148: }
149: ```
150: 
151: Add helpers:
152: 
153: ```go
154: func writeJSON(w http.ResponseWriter, status int, v any) {
155: 	w.Header().Set("Content-Type", "application/json")
156: 	w.WriteHeader(status)
157: 	if err := json.NewEncoder(w).Encode(v); err != nil {
158: 		log.Println(err)
159: 	}
160: }
161: 
162: func writeError(w http.ResponseWriter, status int, message string) {
163: 	writeJSON(w, status, ErrorResponse{Error: message})
164: }
165: ```
166: 
167: Add middleware:
168: 
169: ```go
170: func loggingMiddleware(next http.Handler) http.Handler {
171: 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
172: 		start := time.Now()
173: 
174: 		next.ServeHTTP(w, r)
175: 
176: 		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
177: 	})
178: }
179: ```
180: 
181: Use explicit mux:
182: 
183: ```go
184: mux := http.NewServeMux()
185: mux.HandleFunc("/health", ...)
186: mux.HandleFunc("/expenses", ...)
187: mux.HandleFunc("/expenses/", ...)
188: 
189: log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
190: ```
191: 
192: ## Day 5 Checklist
193: 
194: - All API responses JSON except `/health` can stay plain `OK`.
195: - All errors use `{"error":"..."}` shape.
196: - No `http.Error` left in handlers.
197: - Every error response returns immediately.
198: - Middleware logs each request.
199: - Server uses explicit mux.
200: - `go run .` works from `http-pgsql`.
201: - Curl checks pass:
202: 
203: ```bash
204: curl -i http://localhost:8080/expenses
205: curl -i -X POST http://localhost:8080/expenses -d 'bad'
206: curl -i -X DELETE http://localhost:8080/expenses/not-a-number
207: ```
208: 
209: ## Day 7 Teaching Points
210: 
211: - `go fn()` starts new goroutine.
212: - Goroutines are lightweight threads managed by Go runtime.
213: - `time.NewTicker(interval)` fires events at regular intervals.
214: - `defer ticker.Stop()` prevents resource leak.
215: - `select` waits on multiple channels, picks first ready.
216: - `context.Context` carries cancellation across goroutines.
217: - `signal.NotifyContext` converts OS signals to context cancellation.
218: - `http.Server.Shutdown` stops accepting new requests, drains active ones.
219: - `http.ErrServerClosed` is expected on graceful shutdown, not an error.
220: - Shutdown context must be fresh, not already-canceled signal context.
221: - `<-ctx.Done()` blocks until context is canceled.
222: - `context.WithTimeout` creates deadline for shutdown operation.
223: 
224: ## Teaching Points
225: 
226: - `any` is alias for `interface{}`.
227: - One JSON response path reduces duplicated headers/encoder code.
228: - Structured errors keep clients predictable.
229: - Headers must be set before `WriteHeader`.
230: - Middleware wraps `http.Handler`.
231: - `http.Handler` interface has `ServeHTTP`.
232: - `http.HandlerFunc` adapts a function to handler.
233: - Request flow: server -> middleware -> mux -> handler.
234: - Go package usually maps to folder.
235: - Import path starts with module name from `go.mod`: `expense-tracker`.
236: - `internal/...` packages are importable only from inside the same module tree.
237: - Capitalized identifiers are exported across packages.
238: - Lowercase identifiers are private to package.
239: - Interfaces are satisfied implicitly; no `implements` keyword.
240: - `http.Handler` requires `ServeHTTP(ResponseWriter, *Request)`.
241: - `*http.ServeMux` satisfies `http.Handler`, so `return mux` works from `func NewHandler(...) http.Handler`.
242: - Returning `http.Handler` hides concrete mux implementation from caller.
243: 
244: ## Next Step
245: 
246: Start Day 8: (next topic TBD)

(End of file - total 246 lines)
</content>
```

**Tool: bash**

**Input:**
```json
{
  "command": "cd /home/jyung/Documents/dev/personal/learning/golang && for d in url-shortener http-basic date file; do echo \"=== $d ===\"; rtk ls -la \"$d\" 2>/dev/null; done",
  "description": "List key subdir contents"
}
```

**Output:**
```
=== url-shortener ===
api.http  271B
go.mod  31B
main.go  1.9K
middleware.go  846B
session-ses_18e7.md  7.1K
session-ses_1934.md  6.0K
=== http-basic ===
main.go  1.3K
=== date ===
main.go  118B
=== file ===
read-write-json.go  574B

```

---

## Assistant (Plan · Claude Opus 4.8 · 9.5s)

**Tool: bash**

**Input:**
```json
{
  "command": "cd /home/jyung/Documents/dev/personal/learning/golang/http-goroutine && rtk read go.mod && echo \"=== store ===\" && rtk read internal/store/store.go && echo \"=== response ===\" && rtk read internal/api/response.go",
  "description": "Read module, store, response files"
}
```

**Output:**
```
module expense-tracker

go 1.26.2

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.9.2 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)
=== store ===
package store

import (
	"context"
	"log"
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

func (s *Store) Count(ctx context.Context) (int, error) {
	var count int

	err := s.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM expenses
	`).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
=== response ===
package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
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

```

---

## Assistant (Plan · Claude Opus 4.8 · 5.4s)

**Tool: bash**

**Input:**
```json
{
  "command": "cd /home/jyung/Documents/dev/personal/learning/golang/http-goroutine && rtk read internal/api/handlers.go",
  "description": "Read handlers file"
}
```

**Output:**
```
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

```

---

## Assistant (Plan · Claude Opus 4.8 · 833.4s)

Code understood. Now plan Day 8: Go testing from basic.

## Context

- Day 1-7 done. Expense Tracker API. `http-goroutine` = latest (packages, pgx, goroutines).
- Learner: TS/Node dev, build-first then explain, caveman mode.
- Current tests: only `go test ./...` compile-check. No real tests yet.

## Plan: Day 8 — Go `testing` basics

Go test framework built into stdlib. No Jest/Vitest install. Files end `_test.go`, funcs `func TestXxx(t *testing.T)`, run `go test`.

Build-first progression, simple → real. New folder `testing-basics` to keep clean, isolated from DB-dependent code.

### Step 1 — First test (pure function)
- New folder `testing-basics/` with `go.mod`.
- Write `calc.go` with simple funcs (`Add`, `Divide` returns error).
- Write `calc_test.go`:
  - `func TestAdd(t *testing.T)` basic assert via `if got != want { t.Errorf(...) }`.
  - Show `t.Error` vs `t.Fatal`, `%v` formatting.
- Run `go test`, `go test -v`, `go test -run TestAdd`.

### Step 2 — Table-driven tests
- Convert to Go-idiomatic table pattern: slice of structs `{name, in, want}`, loop, `t.Run(tt.name, ...)` subtests.
- Show why table-driven = Go norm (vs TS `it.each`).
- Run `go test -v` to see subtree output.

### Step 3 — Error-path testing
- Test `Divide` by zero returns error.
- Pattern: check `err != nil`, `errors.Is`, compare messages.

### Step 4 — Coverage + helpers
- `go test -cover`, `go test -coverprofile`, `go tool cover -html`.
- `t.Helper()` for assert helper funcs.

### Step 5 — Test HTTP handler (real project)
- Move to `http-goroutine`, test handler without real DB.
- Introduce `httptest`: `httptest.NewRecorder`, `httptest.NewRequest`.
- Problem surfaces: handlers depend on concrete `*store.Store` (Postgres). Need interface to mock.
- Teach: refactor `NewHandler` to accept interface, write fake store. This connects to dependency-injection / testability (TS DI parallel).

### Teaching points to cover
- `_test.go` convention, same package vs `_test` package (white vs black box).
- `t.Errorf` (continue) vs `t.Fatalf` (stop).
- Table-driven + `t.Run` subtests.
- `t.Helper()`.
- `httptest` recorder/request.
- Interface-based mocking for testability.
- `go test ./...`, `-v`, `-run`, `-cover`, `-race`.

## Questions before build

**Tool: question**

**Input:**
```json
{
  "questions": [
    {
      "question": "Where build Step 1-4 basics?",
      "header": "Location",
      "options": [
        {
          "label": "New testing-basics folder (Recommended)",
          "description": "Clean isolated module for pure-function test learning, then move to http-goroutine for Step 5 real handler tests"
        },
        {
          "label": "Inside http-goroutine",
          "description": "Add test files directly to existing project, skip toy examples"
        }
      ]
    },
    {
      "question": "How far Day 8 go?",
      "header": "Scope",
      "options": [
        {
          "label": "All 5 steps (Recommended)",
          "description": "Basics through real HTTP handler testing with interface mocking"
        },
        {
          "label": "Steps 1-4 only",
          "description": "Pure-function testing fundamentals, save httptest/mocking for Day 9"
        },
        {
          "label": "Steps 1-2 only",
          "description": "Just first test + table-driven, smallest start"
        }
      ]
    },
    {
      "question": "Mock strategy for Step 5 store dependency?",
      "header": "Mocking",
      "options": [
        {
          "label": "Refactor to interface + hand-written fake (Recommended)",
          "description": "Idiomatic Go: define StoreIface, NewHandler takes interface, write fake struct. Teaches testability/DI"
        },
        {
          "label": "Real test DB",
          "description": "Spin Postgres, test against real DB. More setup, integration-style"
        },
        {
          "label": "Decide at Step 5",
          "description": "Reach that point first, choose then"
        }
      ]
    }
  ]
}
```

**Output:**
```
User has answered your questions: "Where build Step 1-4 basics?"="New testing-basics folder (Recommended)", "How far Day 8 go?"="All 5 steps (Recommended)", "Mock strategy for Step 5 store dependency?"="Real test DB". You can now continue with the user's answers in mind.
```

---

## Assistant (Plan · Claude Opus 4.8 · 13.9s)

_Thinking:_

 The user is going with a new folder structure and wants to run through all five steps, using a real Postgres database for the integration tests in Step 5.

Answers locked. Final plan.

## Day 8 — Go `testing` basics (build-first)

### Step 1 — First test, pure function
New folder `testing-basics/` with own `go.mod` (`module testing-basics`).
- `calc.go`: `Add(a, b int) int`, `Divide(a, b int) (int, error)`.
- `calc_test.go`: `TestAdd` using `if got != want { t.Errorf(...) }`.
- Run: `go test`, `go test -v`, `go test -run TestAdd`.

### Step 2 — Table-driven tests
- Rewrite `TestAdd` as table: `[]struct{name string; a,b,want int}`, loop, `t.Run(tt.name, ...)`.
- Idiomatic Go norm. Run `go test -v`, see subtests.

### Step 3 — Error-path
- `TestDivide` covers zero-divisor returns error.
- Check `err != nil`, `errors.Is`, sentinel error `ErrDivideByZero`.

### Step 4 — Coverage + helpers
- `t.Helper()` assert helper.
- `go test -cover`, `-coverprofile=cover.out`, `go tool cover -html=cover.out`.

### Step 5 — Real DB integration test (http-goroutine)
- Add tests in `http-goroutine` hitting actual Postgres (Docker Compose already there).
- `store/store_test.go`: connect via `DATABASE_URL`, test `Add`/`GetAll`/`Delete`/`Count`.
- Pattern: `TestMain(m *testing.M)` for setup/teardown, clean table between tests (`TRUNCATE`).
- `httptest` for handler: `httptest.NewRecorder`, `httptest.NewRequest`, real store wired in.
- Skip if no DB: `t.Skip` when `DATABASE_URL` unreachable / short mode `testing.Short()`.

### Teaching points
- `_test.go` convention, `TestXxx(t *testing.T)`.
- `t.Errorf` (continue) vs `t.Fatalf` (stop).
- Table-driven + `t.Run` subtests.
- `t.Helper()`, sentinel errors + `errors.Is`.
- `httptest` recorder/request.
- `TestMain` setup/teardown, `TRUNCATE` isolation, `testing.Short()`/`-short`.
- Flags: `go test ./...`, `-v`, `-run`, `-cover`, `-race`.

### Verify each step
`go test -v` after every file. Step 5 needs Docker Postgres up.

### PROGRESS.md update
Mark Day 8 done after build, record teaching points + folders.

Ready. Approve to exit plan mode and start building Step 1.

---

