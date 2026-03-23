package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type chunk struct {
	start int
	end   int
}

func main() {
	// generate a large slice of 50000 random items
	data := make([]int, 50000)
	for i := range 50000 {
		data[i] = rand.Intn(100000) + 1
	}

	const numWorkers = 5
	const chunkSize = 10000 // 50000 / 5 workers

	jobs := make(chan chunk, numWorkers)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			workerNode(id, data, jobs)
		}(i)
	}

	// send chunks
	for i := 0; i < len(data); i += chunkSize {
		end := min(i+chunkSize, len(data)) // handle last chunk if slice doesn't divide evenly
		jobs <- chunk{start: i, end: end}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All chunks processed")
}

func workerNode(id int, data []int, jobs <-chan chunk) {
	for c := range jobs {
		fmt.Printf("Worker %d processing indices %d to %d\n", id, c.start, c.end-1)
		processChunk(data[c.start:c.end])
	}
}

func processChunk(chunk []int) {
	sum := 0
	for _, v := range chunk {
		sum += v
	}
	fmt.Printf("chunk sum: %d\n", sum)
}
