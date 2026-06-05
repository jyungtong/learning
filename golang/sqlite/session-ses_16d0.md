# Learning SQLite in Go

**Session ID:** ses_16d0bacf2ffe0X3BScoYxENbSB
**Created:** 6/4/2026, 10:05:46 PM
**Updated:** 6/5/2026, 11:41:42 AM

---

## User

learn sqlite in go

---

## Assistant (Plan · Claude Sonnet 4.6 · 12.1s)



**Tool: read**

**Input:**
```json
{
  "filePath": "/home/jyung/Documents/dev/personal/learning/golang/sqlite"
}
```

**Output:**
```
<path>/home/jyung/Documents/dev/personal/learning/golang/sqlite</path>
<type>directory</type>
<entries>


(0 entries)
</entries>
```

---

## Assistant (Plan · Claude Sonnet 4.6 · 15.0s)

Empty dir. Plan for learning SQLite in Go from scratch.

**Proposed learning structure:**

```
sqlite/
├── go.mod
├── 01_basics/        - open DB, create table, insert, query
├── 02_crud/          - full CRUD with structs
├── 03_prepared/      - prepared statements, transactions
├── 04_migrations/    - schema versioning pattern
└── 05_advanced/      - FTS, JSON columns, WAL mode
```

**Driver choice:**

| Option | Pros | Cons |
|--------|------|------|
| `github.com/mattn/go-sqlite3` | battle-tested, CGo, full SQLite | requires CGo/gcc |
| `modernc.org/sqlite` | pure Go, no CGo | newer, slightly slower |

Recommend `modernc.org/sqlite` — no build toolchain dependency, easier cross-compile.

**Topics per module:**

1. **basics** — `sql.Open`, `db.Exec`, `db.QueryRow`, `db.Query`, `rows.Scan`
2. **crud** — struct mapping, `database/sql` patterns, error handling
3. **prepared** — `db.Prepare`, `tx.Exec`, rollback/commit
4. **migrations** — `PRAGMA user_version`, up/down migration pattern
5. **advanced** — WAL mode, FTS5, JSON1, `PRAGMA` tuning

Questions before I build this out:

1. **CGo OK?** Prefer `mattn/go-sqlite3` (CGo) or `modernc.org/sqlite` (pure Go)?
2. **Scope** — stop at basics/CRUD, or go through advanced topics too?
3. **Style** — standalone runnable examples per topic, or one progressive project (e.g. a small notes app)?

---

## User

start with database/sql

---

## Assistant (Plan · Claude Sonnet 4.6 · 3.5s)

Plan for `database/sql` with SQLite in Go.

**Tool: todowrite**

**Input:**
```json
{
  "todos": [
    {
      "content": "Plan database/sql + SQLite learning structure",
      "status": "in_progress",
      "priority": "high"
    }
  ]
}
```

**Output:**
```
[
  {
    "content": "Plan database/sql + SQLite learning structure",
    "status": "in_progress",
    "priority": "high"
  }
]
```

---

## Assistant (Plan · Claude Sonnet 4.6 · 11.1s)

**Plan: `database/sql` with SQLite in Go**

---

**Module structure:**

```
sqlite/
├── go.mod                  (module: sqlite-learn)
├── 01_connect/
│   └── main.go             open DB, ping, close
├── 02_ddl/
│   └── main.go             CREATE TABLE, DROP TABLE
├── 03_insert/
│   └── main.go             Exec, LastInsertId
├── 04_query/
│   └── main.go             QueryRow + Scan, Query + rows iteration
├── 05_update_delete/
│   └── main.go             Exec with args, RowsAffected
├── 06_transactions/
│   └── main.go             Begin, Commit, Rollback, defer pattern
└── 07_prepared/
    └── main.go             Prepare, Stmt.Exec, Stmt.Query, Stmt.Close
```

**Driver:** `modernc.org/sqlite` (pure Go, no CGo)

**Each file:** standalone `main.go` — run with `go run ./01_connect/`, see output, learn one concept.

**Concepts per module:**

| # | File | `database/sql` API covered |
|---|------|---------------------------|
| 01 | connect | `sql.Open`, `db.Ping`, `db.Close`, `sql.DB` |
| 02 | ddl | `db.Exec` (no args), create/drop table |
| 03 | insert | `db.Exec` with `?` placeholders, `Result.LastInsertId` |
| 04 | query | `db.QueryRow` + `row.Scan`, `db.Query` + `rows.Next/Scan/Close/Err` |
| 05 | update/delete | `db.Exec`, `Result.RowsAffected` |
| 06 | transactions | `db.Begin`, `tx.Exec`, `tx.Commit`, `tx.Rollback`, defer rollback pattern |
| 07 | prepared stmts | `db.Prepare`, `stmt.Exec`, `stmt.Query`, `stmt.Close` |

**Out of scope for now:** connection pool tuning, migrations, advanced pragmas, FTS.

---

Ready to build. Confirm or adjust?

---

## User

use one file to learn all

---

## Assistant (Plan · Claude Sonnet 4.6 · 7.3s)

Single file covering all `database/sql` concepts.

**`main.go` structure:**

```
func main() {
    connectDemo()
    ddlDemo()
    insertDemo()
    queryDemo()
    updateDeleteDemo()
    transactionDemo()
    preparedStmtDemo()
}
```

Each `*Demo()` func is self-contained, prints header + results, teaches one concept. Single in-memory DB (`:memory:`) — no files left on disk.

**File layout:**

```
sqlite/
├── go.mod
└── main.go
```

**All topics in one run:**

| Func | Concept |
|------|---------|
| `connectDemo` | `sql.Open`, `db.Ping`, `db.SetMaxOpenConns` |
| `ddlDemo` | `db.Exec` DDL, error handling |
| `insertDemo` | `?` placeholders, `LastInsertId` |
| `queryDemo` | `QueryRow`+`Scan`, `Query`+`rows.Next/Close/Err` |
| `updateDeleteDemo` | `RowsAffected` |
| `transactionDemo` | `Begin`, `Commit`, `Rollback`, defer pattern |
| `preparedStmtDemo` | `Prepare`, `Stmt` reuse, `stmt.Close` |

Build it?

---

## User

use sqlite

---

## Assistant (Plan · Claude Sonnet 4.6 · 3.8s)

**Tool: todowrite**

**Input:**
```json
{
  "todos": [
    {
      "content": "Init go module + install modernc.org/sqlite",
      "priority": "high",
      "status": "pending"
    },
    {
      "content": "Write main.go with all database/sql demos",
      "priority": "high",
      "status": "pending"
    }
  ]
}
```

**Output:**
```
[
  {
    "content": "Init go module + install modernc.org/sqlite",
    "status": "pending",
    "priority": "high"
  },
  {
    "content": "Write main.go with all database/sql demos",
    "status": "pending",
    "priority": "high"
  }
]
```

---

## Assistant (Plan · Claude Sonnet 4.6 · 10.1s)

Plan:

**Files:**
```
sqlite/
├── go.mod      (module sqlite-learn, go 1.22)
└── main.go
```

**`go.mod` deps:**
- `modernc.org/sqlite` — pure Go SQLite driver, blank import `_ "modernc.org/sqlite"`, driver name `"sqlite"`

**`main.go` structure:**
```go
package main

// imports: database/sql, fmt, log, _ modernc.org/sqlite

func main() {
    db := openDB()         // sql.Open + Ping
    defer db.Close()

    ddlDemo(db)            // CREATE TABLE
    insertDemo(db)         // Exec + LastInsertId
    queryDemo(db)          // QueryRow + Query
    updateDeleteDemo(db)   // Exec + RowsAffected
    transactionDemo(db)    // Begin/Commit/Rollback
    preparedStmtDemo(db)   // Prepare + Stmt reuse
}
```

**DB:** `:memory:` — ephemeral, no cleanup needed.

**Table used throughout:** `users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)`

**Each func:** prints `=== Section Name ===` header, shows results, teaches one concept. No shared state between funcs except the `*sql.DB`.

**Key `database/sql` APIs covered:**

| API | Where |
|-----|-------|
| `sql.Open` | `openDB` |
| `db.Ping` | `openDB` |
| `db.SetMaxOpenConns` | `openDB` |
| `db.Exec` (DDL) | `ddlDemo` |
| `db.Exec` + `Result.LastInsertId` | `insertDemo` |
| `db.QueryRow` + `row.Scan` | `queryDemo` |
| `db.Query` + `rows.Next/Close/Err` | `queryDemo` |
| `Result.RowsAffected` | `updateDeleteDemo` |
| `db.Begin`, `tx.Exec`, `tx.Commit`, `tx.Rollback` | `transactionDemo` |
| `db.Prepare`, `stmt.Exec`, `stmt.Query`, `stmt.Close` | `preparedStmtDemo` |

Ready to implement?

---

## User

show code snippet sql.OPen

---

## Assistant (Plan · Claude Sonnet 4.6 · 6.1s)

```go
// sql.Open validates driver name + DSN format but does NOT connect yet.
// db.Ping() forces actual connection attempt.

db, err := sql.Open("sqlite", ":memory:")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Verify connection is alive
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

Key points:
- `"sqlite"` — driver name registered by `modernc.org/sqlite`
- `":memory:"` — in-memory DB, wiped on `db.Close()`
- `sql.Open` never returns a connection — `Ping` does
- `defer db.Close()` — always pair with Open

---

## User

store as a file

---

## Assistant (Plan · Claude Sonnet 4.6 · 6.5s)

**Plan:**

1. Init go module in `/home/jyung/Documents/dev/personal/learning/golang/sqlite/`
2. Install `modernc.org/sqlite`
3. Write `main.go` with all demo funcs

**File:** `main.go` ~120 lines covering:
- `openDB()` — `sql.Open` + `Ping` + `SetMaxOpenConns`
- `ddlDemo()` — `CREATE TABLE users`
- `insertDemo()` — `Exec` + `LastInsertId`
- `queryDemo()` — `QueryRow`/`Query` + `Scan`
- `updateDeleteDemo()` — `Exec` + `RowsAffected`
- `transactionDemo()` — `Begin`/`Commit`/`Rollback` defer pattern
- `preparedStmtDemo()` — `Prepare` + `Stmt` reuse

**DSN:** `:memory:` (no file on disk, ephemeral)

One question: use a real file-based DB (e.g. `./learn.db`) instead of `:memory:`? File-based lets you inspect state with `sqlite3` CLI between runs.

---

## User

yes, users.db

---

## Assistant (Plan · Claude Sonnet 4.6 · 16.4s)

**Plan: `main.go` + `go.mod` with `users.db`**

---

**Files to create:**

```
sqlite/
├── go.mod
└── main.go
```

---

**`go.mod`:**
```
module sqlite-learn

go 1.22

require modernc.org/sqlite v1.34.5
```

---

**`main.go` structure:**

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "modernc.org/sqlite"
)

const dsn = "users.db"

func main() {
    db := openDB()
    defer db.Close()

    ddlDemo(db)
    insertDemo(db)
    queryDemo(db)
    updateDeleteDemo(db)
    transactionDemo(db)
    preparedStmtDemo(db)
}
```

**Each func — what it teaches:**

| Func | Concept | Key API |
|------|---------|---------|
| `openDB` | connect + verify | `sql.Open`, `db.Ping`, `db.SetMaxOpenConns` |
| `ddlDemo` | schema creation | `db.Exec` DDL |
| `insertDemo` | insert rows | `db.Exec` + `Result.LastInsertId` |
| `queryDemo` | read rows | `db.QueryRow`+`Scan`, `db.Query`+`rows.Next/Close/Err` |
| `updateDeleteDemo` | mutate rows | `db.Exec` + `Result.RowsAffected` |
| `transactionDemo` | atomicity | `db.Begin`, `tx.Exec`, `tx.Commit`, defer `tx.Rollback` |
| `preparedStmtDemo` | stmt reuse | `db.Prepare`, `stmt.Exec`, `stmt.Query`, `stmt.Close` |

**Table:** `users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, age INTEGER NOT NULL)`

**DSN:** `"users.db"` — creates file in working dir, inspectable with `sqlite3 users.db`

**Run:**
```bash
go run .
sqlite3 users.db ".tables"
sqlite3 users.db "SELECT * FROM users;"
```

---

Ready to implement?

---

## User

openDB code snippets

---

## Assistant (Plan · Claude Sonnet 4.6 · 6.4s)

```go
func openDB() *sql.DB {
    // sql.Open validates args but does NOT connect
    db, err := sql.Open("sqlite", "users.db")
    if err != nil {
        log.Fatal(err)
    }

    // SQLite is file-based — 1 writer at a time
    db.SetMaxOpenConns(1)

    // Forces actual connection + file creation
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("=== openDB ===")
    fmt.Println("connected to users.db")
    return db
}
```

Key points:
- `SetMaxOpenConns(1)` — SQLite allows only one writer; prevents `SQLITE_BUSY`
- `Ping()` — creates `users.db` on disk if not exists
- `log.Fatal` — prints error + exits; appropriate for setup failures
- caller owns `db`, responsible for `defer db.Close()`

---

## User

ddl() code snippets

---

## Assistant (Plan · Claude Sonnet 4.6 · 6.3s)

```go
func ddlDemo(db *sql.DB) {
    // db.Exec — for statements that return no rows (DDL, INSERT, UPDATE, DELETE)
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id   INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT    NOT NULL,
            age  INTEGER NOT NULL
        )
    `)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("=== ddlDemo ===")
    fmt.Println("table users created (or already exists)")
}
```

Key points:
- `db.Exec` — use for DDL; returns `(Result, error)`, discard `Result` with `_`
- `CREATE TABLE IF NOT EXISTS` — idempotent; safe to re-run on existing DB
- `AUTOINCREMENT` — SQLite guarantees monotonically increasing IDs (never reuses deleted IDs)
- `NOT NULL` — enforce at DB level, not just app level

---

## User

insertDemo

---

## Assistant (Plan · Claude Sonnet 4.6 · 11.2s)

```go
func insertDemo(db *sql.DB) {
    fmt.Println("=== insertDemo ===")

    // Single insert — ? placeholder prevents SQL injection
    res, err := db.Exec(
        "INSERT INTO users (name, age) VALUES (?, ?)",
        "Alice", 30,
    )
    if err != nil {
        log.Fatal(err)
    }

    id, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("inserted id=%d\n", id)

    // Multiple inserts
    rows := []struct {
        name string
        age  int
    }{
        {"Bob", 25},
        {"Carol", 35},
        {"Dave", 28},
    }

    for _, r := range rows {
        res, err := db.Exec(
            "INSERT INTO users (name, age) VALUES (?, ?)",
            r.name, r.age,
        )
        if err != nil {
            log.Fatal(err)
        }
        id, _ := res.LastInsertId()
        fmt.Printf("inserted id=%d name=%s age=%d\n", id, r.name, r.age)
    }
}
```

Key points:
- `?` placeholders — driver escapes values; never interpolate user input into SQL strings
- `Result.LastInsertId()` — returns the `AUTOINCREMENT` id of inserted row
- `Result.RowsAffected()` — not shown here, covered in `updateDeleteDemo`
- each `db.Exec` is its own implicit transaction; for bulk inserts use an explicit transaction (faster)

---

## User

queryDemo

---

## Assistant (Plan · Claude Sonnet 4.6 · 8.8s)

```go
func queryDemo(db *sql.DB) {
    fmt.Println("=== queryDemo ===")

    // QueryRow — fetch exactly one row
    // use when you expect 0 or 1 result
    var id int
    var name string
    var age int

    err := db.QueryRow(
        "SELECT id, name, age FROM users WHERE name = ?", "Alice",
    ).Scan(&id, &name, &age)
    switch {
    case err == sql.ErrNoRows:
        fmt.Println("alice not found")
    case err != nil:
        log.Fatal(err)
    default:
        fmt.Printf("QueryRow: id=%d name=%s age=%d\n", id, name, age)
    }

    // Query — fetch multiple rows
    // always Close rows, even on error
    rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Query all users:")
    for rows.Next() {
        if err := rows.Scan(&id, &name, &age); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("  id=%d name=%s age=%d\n", id, name, age)
    }

    // Check for iteration error — distinct from query error
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
```

Key points:
- `QueryRow` — never returns error directly; error surfaces on `Scan`; `sql.ErrNoRows` when no match
- `Query` — returns `*sql.Rows`; must `defer rows.Close()` to release connection back to pool
- `rows.Next()` — advances cursor; returns `false` on exhaustion or error
- `rows.Err()` — check after loop; catches mid-iteration errors (network, decode failures)
- `Scan` — args must be pointers, match column order in `SELECT`

---

## User

can we Scan into a struct var?

---

## Assistant (Plan · Claude Sonnet 4.6 · 9.6s)

Yes, but `database/sql` has no built-in struct scanning. Two options:

**Option 1: Manual scan (stdlib only)**
```go
type User struct {
    ID   int
    Name string
    Age  int
}

var u User
err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", 1).
    Scan(&u.ID, &u.Name, &u.Age)
```
Works, but column order must match `Scan` args manually.

---

**Option 2: `sqlx` — struct scanning via reflection**
```go
import "github.com/jmoiron/sqlx"

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
}

var u User
err := db.Get(&u, "SELECT * FROM users WHERE id = ?", 1)

var users []User
err = db.Select(&users, "SELECT * FROM users ORDER BY id")
```

`db` tags map columns to fields — order doesn't matter.

---

| | stdlib `database/sql` | `sqlx` |
|--|--|--|
| struct scan | manual, field-by-field | automatic via `db` tags |
| extra dep | no | `github.com/jmoiron/sqlx` |
| `Get` / `Select` | no | yes |
| wraps `database/sql` | — | yes, fully compatible |

---

For learning `database/sql` internals: stick with manual scan first.
Add `sqlx` after — it's a thin wrapper, easy to layer on.

Want to add `sqlx` examples to the file, or keep it pure stdlib?

---

