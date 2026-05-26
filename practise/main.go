package main

import (
	"fmt"
)

func main() {
	num := 0
	fmt.Println("Enter the number to find frequency in an array : ")
	// fmt.Scan(&num)
	arr := []int{0,1,1,1,5,5,2,2,4,4}
	freq := findFrequencyB(arr, num)
	fmt.Printf("Brute Frequency of a given number %d, exists %d times\n", num, freq)
	num = 6
	arr = []int{1,2,3,4,5,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,7,7,7,8,8,8,9}
	freqO := findFrequencyO(arr, num)
	fmt.Printf("Optimized Frequency of a given number %d, exists %d times\n", num, freqO)
}

//can be applied in any of the array
func findFrequencyB(arr []int, num int) int {
	f := 0
	for i := 0 ; i<len(arr); i++ {
		if arr[i] == num {
			f +=1
		}
	}
	return f
}

//due to linear scan, only O(n) time complexity and O(1) space complexity

//applied in only sorted array
func findFrequencyO(arr []int, num int) int {
	
	firstIndex := findFirst(arr, num)
	if firstIndex == -1 {
		return 0
	}
	lastIndex := findLast(arr, num)

	return lastIndex-firstIndex+1
}

func findFirst(arr []int, num int) int {
	res := -1
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l+(r-l)/2
		if arr[mid] == num {
			res = mid
			r = mid-1
		}else if arr[mid] < num {
			l = mid +1
		}else {
			r = mid -1
		}
	}

	return res
}

func findLast(arr []int, num int) int {
	res := 0
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l+(r-l)/2
		if arr[mid] == num {
			res = mid
			l = mid + 1
		}else if arr[mid] < num {
			l = mid +1
		}else {
			r = mid -1
		}
	}
	return res
}