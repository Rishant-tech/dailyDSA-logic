package main

import (
	"fmt"
)

func main() {
	fmt.Println("Check Even or Odd (5):", CheckEvenOrOdd(5))
	fmt.Println("Closest Number ([1, 3, 7, 8], 5):", ClosestNumber([]int{1, 3, 7, 8}, 5))
	fmt.Println("Dice Probability (6 sides):", DiceProbability(6))
	fmt.Println("Multiplication Table (3):", MultiplicationTable(3))
	fmt.Println("Sum of Naturals (10):", SumOfNaturals(10))
	fmt.Println("Sum of Squares (5):", SumOfSquares(5))
	a, b := SwapTwoNumbers(10, 20)
	fmt.Println("Swap Two Numbers (10, 20):", a, b)
}
