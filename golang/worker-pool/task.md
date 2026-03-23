# Worker Pool Problems

## Problem 1 — Basic Worker Pool with Context Cancellation

Build a worker pool with **5 workers** and **30 jobs**.
- Use a buffered channel to distribute jobs across workers
- Each worker should process jobs concurrently
- Cancel all workers midway (after job 15) using `context.WithCancel`
- Workers must stop gracefully when context is cancelled
- Use `sync.WaitGroup` to wait for all workers to finish

---

## Problem 2 — Worker Pool with Error Handling

Extend the worker pool to handle errors.
- If a worker encounters an error, it should send the error to an `errs` channel
- On first error, cancel all remaining workers using `cancel()`
- Collect and print all errors after all workers finish
- Simulate an error when job number is a multiple of 7
