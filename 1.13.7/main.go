package main

import (
	"fmt"
)

func main() {
	var count, num, min, minCount int
	fmt.Scan(&count)
	fmt.Scan(&min)
	minCount = 1

	for i := 2; i <= count; i++ {
		fmt.Scan(&num)

		if num < min {
			min = num
			minCount = 1
		} else if num == min {
			minCount++
		}
	}

	fmt.Print(minCount)
}
