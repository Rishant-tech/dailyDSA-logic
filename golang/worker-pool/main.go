package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	jobs := make(chan int, 30)
	errs := make(chan error, 30)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			workerNode(ctx, cancel, id, jobs, errs)
		}(i)
	}

	for i := 1; i <= 30; i++ {
		if i == 15 {
			cancel() // cancel context midway
			break
		}
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(errs)

	for err := range errs {
		fmt.Println("Error:", err)
	}

	fmt.Println("All workers have finished their work")
}

func workerNode(ctx context.Context, cancel context.CancelFunc, id int, jobs <-chan int, errs chan<- error) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stopping\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			if err := processJob(id, job); err != nil {
				errs <- err
				cancel() // stop all workers on first error
				return
			}
		}
	}
}

func processJob(id, job int) error {
	fmt.Printf("Worker %d is processing job %d\n", id, job)
	if job%7 == 0 { // simulate error on multiples of 7
		return fmt.Errorf("worker %d failed on job %d", id, job)
	}
	return nil
}
