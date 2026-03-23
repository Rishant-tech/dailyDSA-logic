package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

func superiorWorkerPool() {
	var wg sync.WaitGroup

	jobs := make(chan int, 30)
	errs := make(chan error, 30)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var nextID atomic.Int32
	nextID.Store(5)

	var startWorker func(id int)
	startWorker = func(id int) {
		wg.Add(1)
		go func() {
			err := superiorWorkerNode(ctx, id, jobs)
			if err != nil {
				errs <- err
				newID := int(nextID.Add(1)) - 1
				fmt.Printf("Worker %d failed, restarting as Worker %d\n", id, newID)
				startWorker(newID) // wg.Add(1) called before wg.Done() below
			}
			wg.Done()
		}()
	}

	for i := 0; i < 5; i++ {
		startWorker(i)
	}

	for i := 1; i <= 30; i++ {
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

func superiorWorkerNode(ctx context.Context, id int, jobs <-chan int) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stopping\n", id)
			return nil
		case job, ok := <-jobs:
			if !ok {
				return nil
			}
			if err := processJob(id, job); err != nil {
				return err // stop this worker, supervisor will restart it
			}
		}
	}
}
