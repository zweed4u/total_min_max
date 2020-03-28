package main

import (
	"fmt"
	"math"
)

func miniMaxSum(arr []int) string {
	min := math.MaxInt32
	max := 0
	total := 0
	for _, number := range arr {
		min = int(math.Min(float64(min), float64(number)))
		max = int(math.Max(float64(max), float64(number)))
		total += number
	}
	return fmt.Sprintf("%d %d", total-max, total-min)
}

func main() {
	s := miniMaxSum([]int{1, 2, 3, 4})
	fmt.Println(s)
}
