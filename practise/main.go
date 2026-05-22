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

// 	waterTrapped := trapRainWaterBrute(arr)

// 	fmt.Println("Total water trapped:", waterTrapped)

// 	waterTrapped = trapRainWaterOptimalTwoPointer(arr)

// 	fmt.Println("Total water trapped:", waterTrapped)
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

func main () {
	arr := []int{1,2,3,3,4,4,5,5,6,6,7,8,9}
	var n int

	fmt.Println("Enter the number you want to find frequency for : ")

	fmt.Scan(&n)

	frequency := findFrequency(arr,n)

	fmt.Printf("Frequency of the number %d, is : %d\n",n,frequency)
}

func findFrequency(arr []int, num int) int {
	freq := 0

	for i :=0; i<len(arr); i++{
		if arr[i] == num {
			freq++
		}
	}
	return freq
}