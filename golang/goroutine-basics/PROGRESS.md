# Goroutine Basics Progress

Goal:
Learn goroutines from mechanics to backend patterns before returning to Expense Tracker Day 7.

Run style:
Each future example file lives directly in this folder.

Example command:
go run 01-blocking.go
go run 15-race-detector.go
go run -race 15-race-detector.go

## Phase 1: Mechanics

- [ ] `01-blocking.go` - normal function call blocks
- [ ] `02-goroutine.go` - `go fn()` starts concurrent work
- [ ] `03-sleep-wait.go` - crude wait with `time.Sleep`
- [ ] `04-waitgroup.go` - proper wait with `sync.WaitGroup`
- [ ] `05-multiple.go` - multiple goroutines, nondeterministic order
- [ ] `06-channel-unbuffered.go` - unbuffered send waits for receive
- [ ] `07-channel-buffered.go` - buffered send waits only when buffer full
- [ ] `08-channel-close-range.go` - sender closes, receiver ranges
- [ ] `09-directional.go` - `chan<-` and `<-chan` compile-time safety
- [ ] `10-select.go` - wait on multiple channel operations

## Phase 2: Safety

- [ ] `11-context-cancel.go` - `ctx.Done()` stops goroutine
- [ ] `12-context-timeout.go` - `context.WithTimeout` sets deadline
- [ ] `13-ticker.go` - repeated work with `time.Ticker`
- [ ] `14-goroutine-leak.go` - goroutine stuck forever
- [ ] `15-race-detector.go` - intentional race, use `go run -race`
- [ ] `16-mutex-fix.go` - fix shared state race with `sync.Mutex`

## Phase 3: Patterns

- [ ] `17-worker-pool.go` - fixed workers, job channel, result channel
- [ ] `18-job-queue-timeout.go` - timeout while waiting for job/result
- [ ] `19-rate-limiter.go` - channel token bucket
- [ ] `20-fan-out-fan-in.go` - split work, merge results
- [ ] `21-graceful-shutdown.go` - signal, context, WaitGroup, drain

## Phase 4: Backend Mapping

- [ ] `22-async-report.go` - background expense report sketch
- [ ] `23-request-timeout.go` - middleware timeout sketch
- [ ] `24-background-cleanup.go` - periodic cleanup job sketch
- [ ] `25-day7-mapping.go` - map basics back to HTTP goroutine project

## Mental Models

- Goroutine = lightweight concurrent function.
- `go fn()` starts work and returns immediately.
- `main` exiting kills remaining goroutines.
- `time.Sleep` is demo-only waiting.
- `WaitGroup` waits for known goroutines.
- Channel sends/receives can block.
- Sender owns channel close.
- `select` chooses first ready case.
- Context tells goroutines when to stop.
- Every goroutine needs an exit path.
- Shared memory needs protection: mutex, channel ownership, or atomic.

## Common Mistakes

- Starting goroutine but not waiting.
- Closing channel from receiver side.
- Sending on closed channel.
- Ranging over channel nobody closes.
- Assuming goroutine output order.
- Using `Sleep` as production coordination.
- Ignoring `ctx.Done()`.
- Creating goroutine leak.
- Writing shared variable from many goroutines without mutex.
- Fire-and-forget goroutine inside HTTP handler without lifecycle owner.

## Bridge To Expense Tracker

- Day 7 server goroutine maps to long-running background task.
- Day 7 ticker maps to repeated background job.
- Day 7 `ctx.Done()` maps to shutdown signal.
- Day 7 `server.Shutdown()` maps to graceful stop/drain.
- Future worker pool maps to async backend jobs.
- Future rate limiter maps to API throttling.
- Future request timeout maps to middleware and DB query deadlines.
