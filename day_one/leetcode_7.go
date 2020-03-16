package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(345))
}

func reverse(x int) int {
	var min = math.MinInt32
	var max = math.MaxInt32
	var y int = 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
		if y < min || y > max {
			return 0
		}
	}
	return y
}
