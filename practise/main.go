package main

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
EXAMPLE

Input:  height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: 6 units of water are trapped between the bars.
GOOGLE APPROACH

For each index, water held = min(maxLeft, maxRight) - height[i]
Brute force O(n²): for each bar scan left and right maxima
Optimized O(n): precompute prefix maxLeft[] and suffix maxRight[] arrays
Best: two-pointer O(n) time O(1) space — shrink from both ends
Google expects the two-pointer approach with clear explanation of invariant

*/

//brute force method

import "fmt"

// func main() {

// 	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}

// 	waterTrapped := trapoptimized(arr)

// 	fmt.Println("Total water trapped:", waterTrapped)

// 	// waterTrapped = trapRainWaterOptimalTwoPointer(arr)

// 	// fmt.Println("Total water trapped:", waterTrapped)
// }

func trapRainWaterBrute(arr []int) int {

	totalWater := 0

	for i:=1;i<len(arr)-1;i++ {

		maxLeft:=0
		maxRight:=0

		for j:=0;j<i;j++ {

			if arr[j]>maxLeft{
				maxLeft=arr[j]
			}
		}

		for j:=i+1;j<len(arr);j++ {

			if arr[j]>maxRight{
				maxRight=arr[j]
			}
		}

		water:=minValue(maxLeft,maxRight)-arr[i]
		if water>0{
			totalWater+=water
		}
	}

	return totalWater
}

func minValue(a,b int) int {

	if a<b{
		return a
	}

	return b
}

func trapRainWaterOptimalTwoPointer(arr []int) int {

	l,r := 0,len(arr)-1
	lmax, rmax, totalWater := 0,0,0

	for l<r {
		fmt.Printf("value %d , %d \n",l,r)

		if arr[l] < arr[r]{
			if arr[l] >= lmax {
				lmax = arr[l]
			}else{
				totalWater = totalWater+lmax-arr[l]
			}
			l++
		}else{
			if arr[r]>=rmax{
				rmax = arr[r]
			}else{
				totalWater = totalWater + rmax -arr[r]
			}
			r--
		}
	}

	fmt.Printf("l : %d, r: %d : \n", lmax, rmax)

	fmt.Println("values : ", arr)

	return totalWater
}

//frequency of number in sorted array
// func main () {
// 	arr := []int{1,2,3,3,4,4,5,5,6,6,7,8,9}
// 	var n int

// 	fmt.Println("Enter the number you want to find frequency for : ")

// 	fmt.Scan(&n)

// 	frequency := findFrequencyBrute(arr,n)

// 	fmt.Printf("Frequency of the number %d, is : %d\n",n,frequency)

// 	frequency = findFrequencyBinarySearch(arr,n)

// 	fmt.Printf("Frequency of the number %d, is : %d\n",n,frequency)
// }

func findFrequencyBrute(arr []int, num int) int {
	freq := 0

	for i :=0; i<len(arr); i++{
		if arr[i] == num {
			freq++
		}
	}
	return freq
}

func findFrequencyBinarySearch(arr []int, num int) int {
	firstIndex := findFirstOccurrence(arr, num)
	lastIndex := findLastOccurrence(arr, num)
	var frequency int

	if firstIndex == -1 {
		fmt.Printf("The number %d does not exist in the array.\n", num)
	} else {
		frequency := lastIndex - firstIndex + 1
		fmt.Printf("Frequency of the number %d is: %d\n", num, frequency)
	}

	return frequency
}

func findFirstOccurrence(arr []int, num int) int{
	left, right := 0, len(arr)-1
	result := -1

	for left <=right {

		mid := left+(right-left)/2
		if arr[mid] == num {
			result = mid
			right = mid -1 
		}else if arr[mid] < num {
			left = mid+1
		}else {
			right = mid -1
		}

	}

	return result
}

func findLastOccurrence(arr []int, num int) int{
	left, right := 0, len(arr)-1
	result := -1

	for left <= right{
		mid := left + (right-left)/2

		if arr[mid] == num{
			result = mid
			left = mid + 1
		}else if arr[mid]<num{
			left = mid+1
		}else{
			right = mid - 1
		}

	}

	return result
}

func trap(arr []int) int{
	//brute force method
	totalWater := 0

	for i := 0; i <= len(arr)-1; i++ {
		lMax, rMax := 0,0

		for j:=0; j<i;j++ {
			if arr[j] > lMax {
				lMax = arr[j]
			}
		}

		for j:=i; j<len(arr); j++ {
			if arr[j] > rMax{
				rMax = arr[j]
			}
		}
		water := minWater(lMax, rMax)-arr[i]
		if water>0{
			totalWater+=water
		}
	}

	return totalWater
}

func minWater(a,b int) int{
	if a<b {
		return a
	}
	return b
}


func trapoptimized(arr []int) int {
	lMax, rMax,totalWater := 0,0,0

	l,r := 0, len(arr)-1

	for l<r {

		if arr[l]<arr[r] {
			if arr[l] > lMax {
				lMax = arr[l]
			}else{
				totalWater += lMax - arr[l]
			}
			l++
		}else {
			if arr[r] > rMax{
				rMax = arr[r]
			}else {
				totalWater += rMax - arr[r]
			}
			r--
		}

	}

	return totalWater 
}



func main() {
	arr := []int{1,2,2,3,3,4,4,4,4,5,5,6,7,8,9,9,9,9}
	var num int
	fmt.Print("Hello Please enter the value to find frequency : ")
	fmt.Scan(&num)
	freq := findFrequencyBrutee(arr,num)

	fmt.Printf("The frequncy of number %d , is exists %d times\n", num, freq)
}

//brute force
func findFrequencyBrutee(arr []int, num int) int {
	freq := 0

	//do linear scan

	for i:=0 ; i<len(arr) ; i++ {
		if arr[i] == num{
			freq++
		}
	}
	return freq
}