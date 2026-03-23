package main

import "fmt"

func main() {
	for {
		var choice int
		fmt.Println("1: normal pool | 2: superior pool | 3: both | 0: exit")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			normalWorkerPool()
		case 2:
			superiorWorkerPool()
		case 3:
			normalWorkerPool()
			superiorWorkerPool()
		default:
			fmt.Println("Invalid choice")
		}
	}
}
