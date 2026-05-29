# Expired URL cleanup design review

**Session ID:** ses_18e797b77ffeJYgy5T09f2gG2r
**Created:** 5/29/2026, 10:18:21 AM
**Updated:** 5/29/2026, 10:32:11 AM

---

## User

@session-ses_1934.md learn background cleanup. 
my idea is change `store` to map[string]UrlBody{}

UrlBody struct with {Url: string, Expired_at: string, Is_expired: bool}

start a go routine in main, the function use time.ticker to receive 1sec channel, use select case time <-Time.ticker(), loop every maps' objects and check if current time more than expired_at, if yes then mark is_expired to true.

challenge my design

---

## Assistant (Plan · Deepseek V4 Pro · 19.3s)

_Thinking:_

The user wants me to review and challenge their design for background cleanup of expired URLs in their URL shortener. Let me read their current code first, then analyze their proposed design.

Their current code:
- `store` is `map[string]string{}` — maps short code to URL
- They want to change to `map[string]UrlBody{}` where `UrlBody` has `{Url, Expired_at, Is_expired}`
- Start a goroutine with `time.Ticker` (1 second) that loops all map entries and marks expired ones

Let me think about the design challenges/critiques.

Design review (GOOD ideas, but optimizations needed):

**1. Tick interval too aggressive**
1-second tick for a map that might have thousands of entries is wasteful. For URL expiry, 10-30 seconds is fine.

**2. O(n) scan every tick is brute force**
Scanning ALL entries every tick is O(n) CPU burn. Better: **priority queue (heap)** sorted by `Expired_at`. Only check the next-to-expire entry each tick. When its time hasn't come yet, nothing else has either.

```go
type ExpiryHeap []UrlBody // implements heap.Interface, sorted by Expired_at
```

On tick: peek top of heap. If top not expired → skip whole tick. If expired → pop and mark.

**3. `Is_expired` boolean is redundant**
If `Expired_at` exists, you can derive expiry: `time.Now().After(expiredAt)`. Boolean adds state you must synchronize and risks inconsistency. Check at read-time instead:

```go
func (u *UrlBody) IsExpired() bool {
    return time.Now().After(u.ExpiredAt)
}
```

**4. `Expired_at` as `string` is fragile**
Use `time.Time`. Parse on input, compare natively. String comparisons are error-prone with timezones/formats.

**5. Mutex contention: cleanup locks same mutex as reads/writes**
If cleanup holds write lock while iterating all entries, reads block. Solutions:
- Cleanup acquires lock, collects expired keys, releases lock, then deletes (two-phase)
- Or use `sync.Map` for read-heavy workloads

**6. Missing: actual deletion**
Your plan marks `Is_expired = true` but never deletes. Map grows forever. Decide: mark-and-filter on read, or actually delete. If deleting, ensure cleanup removes entries not just marks them.

**7. Consider `context.Context` for graceful shutdown**
Goroutine should receive a `context.Context` or a `chan struct{}` so it can stop cleanly when server shuts down (no orphan goroutines).

---

**Revised structure:**

```go
type UrlBody struct {
    URL       string
    ExpiredAt time.Time
}

func (u UrlBody) IsExpired() bool {
    return time.Now().After(u.ExpiredAt)
}

// Heap-based cleanup
type expiryHeap []struct {
    code     string
    expiredAt time.Time
}
```

**Cleanup goroutine pattern:**
```go
func cleanupExpired(ctx context.Context, h *expiryHeap, store *map[string]UrlBody, mu *sync.RWMutex) {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            mu.Lock()
            // Pop all expired from heap, delete from store
            mu.Unlock()
        }
    }
}
```

**Core challenge:** Heap requires knowing expiry time at insert. When you shorten, compute `time.Now().Add(ttl)` and push to heap. Heap lets you skip full scans.

Want me to implement this design?

---

## User

i'm still learning, suggest me learning path from simple and interate later

---

## Assistant (Plan · Claude Opus 4.8 · 21.1s)

Good instinct. Learn by iterating. Start simplest version of YOUR idea, then improve once it works. Here path:

## Stage 0: current baseline
`store = map[string]string{}`. Working. Don't break it.

## Stage 1: your original idea (simplest, do FIRST)
Goal: learn goroutine + `time.Ticker` + `select`. Accept inefficiency.

- Change `UrlBody` struct: `{URL string, ExpiredAt time.Time, IsExpired bool}`
  - Use `time.Time` not `string`. One small upgrade. Comparing times native = easy.
- `store = map[string]UrlBody{}`
- On shorten: set `ExpiredAt = time.Now().Add(30 * time.Second)` (short, so you SEE it expire while testing)
- Goroutine in `main`:
```go
go func() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            mu.Lock()
            for code, body := range store {
                if time.Now().After(body.ExpiredAt) {
                    body.IsExpired = true
                    store[code] = body // map values not addressable, must reassign
                }
            }
            mu.Unlock()
        }
    }
}()
```
- Read handler: check `if body.IsExpired { 404 }`

**Learning here:** goroutine lifecycle, ticker, channel receive in select, mutex with background writer, map value reassignment gotcha.

This is exactly your design. Ship it. Get it working.

## Stage 2: delete instead of mark
Problem you'll notice: map grows forever, `IsExpired=true` entries never removed.

- In cleanup loop, `delete(store, code)` instead of marking
- Drop `IsExpired` bool entirely — if not in map, it's gone = 404 naturally
- Read handler simplifies: `body, ok := store[code]; if !ok { 404 }`

**Learning:** why redundant state is bad, `delete()` on map during iteration (safe in Go).

## Stage 3: derive expiry at read (lazy)
Question to sit with: do you even NEED the goroutine?

- On read: `if time.Now().After(body.ExpiredAt) { delete + 404 }`
- Goroutine becomes optional — only reclaims memory for links never read again

**Learning:** lazy vs eager cleanup tradeoff. Real systems use both.

## Stage 4: graceful shutdown
Your goroutine runs forever — never stops cleanly.

- Pass `context.Context`, add `case <-ctx.Done(): return` to select
- Wire `signal.Notify` + `server.Shutdown(ctx)` in main

**Learning:** context cancellation, no orphan goroutines, graceful HTTP shutdown.

## Stage 5: optimize scan (ONLY if you want)
Now revisit my heap critique. O(n) scan fine for thousands. Heap matters at scale.

- Min-heap by `ExpiredAt`, peek top, skip tick if top not expired
- `container/heap` package

**Learning:** `heap.Interface`, amortized cost, when optimization is worth complexity.

---

## My advice
Do **Stage 1 exactly as you designed it**. My earlier critique was all valid but premature — optimizing before it works is the trap. Build the naive version, FEEL the problems (map grows, no shutdown), THEN fix them. That's how the lessons stick.

One non-negotiable upgrade even at Stage 1: `time.Time` over `string` for `ExpiredAt`. String dates = pain, no learning value.

Want me to start implementing Stage 1? (toggle out of plan mode first)

---

