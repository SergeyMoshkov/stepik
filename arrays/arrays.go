package main

import "fmt"

func main() {

	workArray := [10]uint8{}
	outputStr := ""
	for i := 0; i < 10; i++ {
		var number uint8
		fmt.Scan(&number)
		workArray[i] = number
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 1; j++ {
			var a uint8
			var b uint8
			fmt.Scan(&a, &b)
			workArray[a], workArray[b] = workArray[b], workArray[a]
			// outputStr += fmt.Sprintf("%d ", "%d ", b, a)

		}
	}

	for _, v := range workArray {
		outputStr += fmt.Sprintf("%d ", v)
	}
	fmt.Print(outputStr)
}
