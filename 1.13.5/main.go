package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	c := float64(a + b) / 2
	fmt.Print(c)
}
