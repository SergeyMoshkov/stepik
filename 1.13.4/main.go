package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный.")
	} else {
		fmt.Println("Непрямоугольный.")
	}
}
