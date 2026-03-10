package main

func ClosestNumber(arr []int, target int) int {
	closest := arr[0]
	for _, num := range arr {
		if abs(num-target) < abs(closest-target) {
			closest = num
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
