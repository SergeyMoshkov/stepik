package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for number >= 10 {
		sum := 0
		for number > 0 {
			digit := number % 10
			sum += digit

			number /= 10
			
		}

		number = sum
		fmt.Println(number)
	}

	fmt.Print(number)
}

