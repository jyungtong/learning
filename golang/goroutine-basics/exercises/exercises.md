# Real-World Go Exercises

Based on goroutine-basics progress. Categorized by level.

---

## Beginner

These reinforce what you already know with minimal new surface area.

**1. Parallel file hasher**
Given a list of file paths, compute the MD5/SHA256 of each file concurrently using a WaitGroup. Print results as they complete.
- Practices: goroutines, WaitGroup, channels for results

**2. Concurrent URL checker**
Take a slice of URLs, check each with `http.Get` concurrently, report which ones are up/down.
- Practices: goroutines, WaitGroup, collecting results safely

**3. Word frequency counter**
Read a large text file, split into chunks, count word frequency per chunk concurrently, merge results.
- Practices: fan-out/fan-in, mutex or channel-based merge

---

## Intermediate

These require combining multiple patterns you've learned.

**4. HTTP server with per-request timeout**
Build a small `net/http` server. Each handler calls a slow "database" function (simulated with `time.Sleep`). Add a middleware that cancels the request context after 500ms.
- Practices: `context.WithTimeout`, middleware, real HTTP

**5. Worker pool with retry**
Build a worker pool that processes jobs. If a job fails, retry up to 3 times with exponential backoff. Use `errgroup` to propagate the first unrecoverable error.
- Practices: worker pool, errgroup, backoff logic

**6. Rate-limited API client**
Write a client that calls an external API (e.g., GitHub's public API). Limit to N requests/second using a token bucket channel. Handle context cancellation.
- Practices: rate limiter, context, real HTTP client

**7. Pipeline: CSV → transform → write**
Read a CSV file, pass rows through a 3-stage pipeline (parse → validate → write to output file), each stage running concurrently.
- Practices: pipeline pattern, directional channels, graceful shutdown

---

## Advanced

These mirror production backend concerns directly.

**8. Background job scheduler**
Build a scheduler that accepts jobs with a "run at" time. A background goroutine wakes up at the right time and executes each job. Support graceful shutdown that drains pending jobs.
- Practices: ticker, context, WaitGroup, graceful shutdown, real scheduling logic

**9. Connection pool**
Implement a simple connection pool (e.g., for a mock DB). Workers borrow connections, use them, return them. Handle the case where all connections are busy (block or timeout).
- Practices: buffered channels as semaphores, context timeout, resource lifecycle

**10. Pub/sub broker**
Build an in-memory pub/sub system. Publishers send messages to topics. Multiple subscribers receive from their subscribed topics. Support subscriber unsubscribe and graceful shutdown.
- Practices: fan-out, channel management, goroutine lifecycle, mutex for subscriber map

**11. Expense Tracker: async report generation**
This is your direct bridge. In your Day 7 project, add a `POST /reports` endpoint that kicks off a background goroutine to generate a report. The client polls `GET /reports/:id` for status. Use a context with timeout to cancel stalled report jobs.
- Practices: everything — this is the real target

---

## Suggested Order

Start with **#2 (URL checker)** — it's quick, uses real I/O, and you'll see goroutine benefits immediately. Then jump to **#4 (HTTP server with timeout)** since it maps directly to your Expense Tracker work. Save **#11** for when you return to Day 7.
