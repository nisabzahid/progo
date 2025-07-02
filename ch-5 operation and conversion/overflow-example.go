package main

import (
	"fmt"
	"math"
)

func main() {
	var max = math.MaxInt64
	var min = math.MinInt64
	var maxFloat = math.MaxFloat64
	var minFloat = -math.MaxFloat64

	fmt.Println(max * 2) // Overflow, result is negative
	fmt.Println(min * 2) // Underflow, result is positive
	fmt.Println(max + 1) // Overflow, result is negative
	fmt.Println(math.IsInf(maxFloat*2, 0))
	fmt.Println(math.IsInf(minFloat*2, 0)) // Underflow, result is positive
}
