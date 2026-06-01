package main

/*
Problem: Longest Subarray with Sum K

Description:
Given an array of integers arr and an integer k, write a function in Go to find the length of the longest contiguous subarray whose sum equals k.

Example:
arr := []int{1, -1, 5, -2, 3}
k := 3
Output: 4  // The subarray [1, -1, 5, -2] sums to 3

arr := []int{-2, -1, 2, 1}
k := 1
Output: 2  // The subarray [-1, 2] sums to 1

*/

import (
	"fmt"
)

func main() {
	arr := []int{1, -1, 5, -2, 3}
	var num int
	fmt.Println("Enter the num : ")
	fmt.Scan(&num)

	subArrSize, subArr := findSubArr(arr, num)

	fmt.Printf("The size of sub array for sum %d, is %d \n", num, subArrSize)
	fmt.Println("The sub Arr : ", subArr)
}

func findSubArr(arr []int, num int) (int, []int){
	preFix := make(map[int]int)
	sum := 0
	maxLen := 0 
	start := 0

	//{1, -1, 5, -2, 3} 0,1,2,3,4 //3

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum == num { //num 3 true for 3
			maxLen = i+1
		}

		if idx, found := preFix[sum - num]; found {
			lenn := i - idx
			if lenn > maxLen {
				maxLen = lenn
				start += 1
			}
		}
		if _, exists := preFix[sum]; !exists {
			preFix[sum] = i 
		}

	}
	return maxLen, arr[start:maxLen]
	//map - [-2,0], [-1,1], [5,2], 
}