package main

import (
	"fmt"
	"math"
)

func minInitialEnergy(heights []int) int {
	energy := 0
	nextEnergy := 0

	for i := len(heights) - 1; i >= 0; i-- {
		energy = int(math.Ceil((float64(heights[i]) + float64(nextEnergy)) / 2.0))
		nextEnergy = energy
	}

	return energy
}

func main() {
	var N int
	fmt.Scan(&N)

	heights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&heights[i])
	}

	fmt.Println(minInitialEnergy(heights))
}
