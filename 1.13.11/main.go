package main

import "fmt"

func main() {
	var n int
	var s int = 2
	fmt.Scan(&n)

	result := 1
	fmt.Print(1, " ")
	for i := 0; i < n; i++ {
		result *= s
		if result > n {
			break
		}
		fmt.Print(result, " ")
	}
}
