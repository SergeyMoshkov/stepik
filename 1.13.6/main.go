package main

import "fmt"

func main() {
	var count, num, zeroCount int
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Scan(&num)

		if num == 0 {
			zeroCount++
		}
	}

	fmt.Print(zeroCount)
}
