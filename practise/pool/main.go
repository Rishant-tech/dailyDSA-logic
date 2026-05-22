package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		results <- task * 2 // Example processing (doubling the task)
	}
}

func main() {
	numTasks := 50
	numWorkers := 5

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	for t := 1; t <= numTasks; t++ {
		tasks <- t
	}
	close(tasks)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}