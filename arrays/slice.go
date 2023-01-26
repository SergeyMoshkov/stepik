package main

import "fmt"

func Slice() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Базовый массив: %v\n", baseArray)

	baseSlice := baseArray[4:8]
	fmt.Printf(
		"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)
}
