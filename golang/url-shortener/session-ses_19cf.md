# Go learning project ideas

**Session ID:** ses_19cf3c56dffeHxcpxqIQIOsXv7
**Created:** 5/26/2026, 2:50:06 PM
**Updated:** 5/26/2026, 2:53:48 PM

---

## User

suggest some exercise / side project i can built with go for learning purpose

---

## Assistant (Plan · Claude Sonnet 4.6 · 8.6s)

Go learning projects by difficulty:

**Beginner**
- CLI todo app — flags, file I/O, structs
- URL shortener — HTTP server, maps, basic routing
- File watcher — `os`, `filepath`, goroutines

**Intermediate**
- REST API (no framework) — `net/http`, JSON, middleware, error handling
- Web scraper — `net/http`, HTML parsing, concurrency with goroutines/channels
- Rate limiter — channels, tickers, sync primitives
- Key-value store (in-memory) — maps, mutexes, optional persistence

**Advanced**
- TCP chat server — `net`, goroutines, channels, connection management
- Mini job queue — goroutines, channels, worker pools
- Static site generator — file I/O, templates, CLI
- DNS resolver — UDP, binary protocol parsing

**What's your current Go level?** Knowing that + what domain interests you (systems, web, CLI, networking) narrows it down fast.

---

## User

url shortener, what components i can learn 

---

## Assistant (Plan · Claude Sonnet 4.6 · 10.9s)

URL shortener covers a solid slice of Go fundamentals:

**Core Go concepts**

- `net/http` — handler funcs, `ServeMux`, request/response lifecycle
- Structs + methods — modeling your URL store
- Maps + `sync.RWMutex` — concurrent-safe in-memory storage
- Error handling patterns — Go's explicit `error` returns
- `encoding/json` — request/response serialization

**HTTP specifics**
- Route matching (GET `/s/{code}` → redirect, POST `/shorten` → create)
- HTTP status codes — `301`/`302` redirects, `404`, `400`, `409`
- Middleware — logging, request ID injection

**If you extend it**
- `database/sql` + SQLite/Postgres — persistence layer
- `os`/`flag` — CLI config (port, DB path)
- Goroutines — background cleanup of expired links
- Testing — `net/http/httptest` for handler tests

**Stretch goals that teach more**
- Rate limiting per IP — channels or `golang.org/x/time/rate`
- Custom short codes — input validation, collision handling
- Expiring links — time.Time, background goroutine with ticker

Start with in-memory + pure `net/http` (no frameworks). That forces you to understand what frameworks abstract away.

Want a project structure / implementation plan?

---

