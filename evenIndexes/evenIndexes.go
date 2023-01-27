package main

import "fmt"

func main() {
	var N uint
	outputString := ""
	fmt.Scan(&N)
	arr := make([]uint, N)
	
	for idx, val := range arr {
		fmt.Scan(&val)
		if idx % 2 == 0 {
			outputString += fmt.Sprintf("%d ", val) 
		}
	}
	
	fmt.Println(outputString)

}