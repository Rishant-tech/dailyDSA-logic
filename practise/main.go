package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////
// TWO SUM
/////////////////////////////////////////////

func twoSum(numbers []int, target int) []int {

	numberMap := make(map[int]int)

	for index := 0; index < len(numbers); index++ {

		currentValue := numbers[index]

		requiredValue := target - currentValue

		if previousIndex, exists := numberMap[requiredValue]; exists {

			return []int{previousIndex, index}
		}

		numberMap[currentValue] = index
	}

	return []int{}
}

/////////////////////////////////////////////
// WORKER POOL
/////////////////////////////////////////////

func worker(workerId int, jobs <-chan int, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	for job := range jobs {

		fmt.Printf("Worker %d processing %d\n", workerId, job)
	}
}

func workerPoolDemo() {

	fmt.Println("\n===== Worker Pool =====")

	jobs := make(chan int, 10)

	var waitGroup sync.WaitGroup

	numberOfWorkers := 3

	for workerIndex := 1; workerIndex <= numberOfWorkers; workerIndex++ {

		waitGroup.Add(1)

		go worker(workerIndex, jobs, &waitGroup)
	}

	for jobNumber := 1; jobNumber <= 10; jobNumber++ {

		jobs <- jobNumber
	}

	close(jobs)

	waitGroup.Wait()
}

/////////////////////////////////////////////
// PRODUCER CONSUMER
/////////////////////////////////////////////

func producer(queue chan<- int) {

	for value := 1; value <= 5; value++ {

		fmt.Println("Produced:", value)

		queue <- value
	}

	close(queue)
}

func consumer(queue <-chan int, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	for value := range queue {

		fmt.Println("Consumed:", value)
	}
}

func producerConsumerDemo() {

	fmt.Println("\n===== Producer Consumer =====")

	queue := make(chan int, 5)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go producer(queue)

	go consumer(queue, &waitGroup)

	waitGroup.Wait()
}

/////////////////////////////////////////////
// RATE LIMITER
/////////////////////////////////////////////

func rateLimiterDemo() {

	fmt.Println("\n===== Rate Limiter =====")

	limiter := time.Tick(time.Second)

	for request := 1; request <= 5; request++ {

		<-limiter

		fmt.Println("Processing request:", request)
	}
}

/////////////////////////////////////////////
// TOP K ELEMENTS USING HEAP
/////////////////////////////////////////////

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {

	old := *h

	length := len(old)

	item := old[length-1]

	*h = old[:length-1]

	return item
}

func topK(numbers []int, k int) []int {

	minHeap := &MinHeap{}

	heap.Init(minHeap)

	for _, value := range numbers {

		heap.Push(minHeap, value)

		if minHeap.Len() > k {

			heap.Pop(minHeap)
		}
	}

	return *minHeap
}

/////////////////////////////////////////////
// MUTEX / CONCURRENCY
/////////////////////////////////////////////

var counter int

var mutex sync.Mutex

func increment(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	mutex.Lock()

	counter++

	mutex.Unlock()
}

func concurrencyDemo() {

	fmt.Println("\n===== Concurrency =====")

	var waitGroup sync.WaitGroup

	for value := 1; value <= 1000; value++ {

		waitGroup.Add(1)

		go increment(&waitGroup)
	}

	waitGroup.Wait()

	fmt.Println("Counter:", counter)
}

/////////////////////////////////////////////
// MAIN
/////////////////////////////////////////////

func main() {

	fmt.Println("===== Two Sum =====")

	input := []int{2, 7, 11, 15}

	fmt.Println(twoSum(input, 9))

	workerPoolDemo()

	producerConsumerDemo()

	rateLimiterDemo()

	fmt.Println("\n===== Top K =====")

	numbers := []int{5, 9, 1, 4, 10, 7}

	fmt.Println(topK(numbers, 3))

	concurrencyDemo()
}