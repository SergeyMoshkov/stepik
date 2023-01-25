package main

import (
	"fmt"
	"reflect"
)

func main() {

	var workArray [10]uint
	var outputStr string
	var number uint
	var c uint

	// считываем числа с ввода и добавляем в массив
	for i := 0; i < 10; i++ {
		fmt.Scan(&number)
		workArray[i] = number
	}

	// считываем индексы массива с входа и переставляем значения в массиве попарно
	for i := 0; i < 3; i++ {
		var a uint
		var b uint
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	// добавляем значения массива в строку
	for _, v := range workArray {
		outputStr += fmt.Sprintf("%d ", v)
	}

	fmt.Print(outputStr)

	//  проверяем соответствует ли тип значений массива заявленому типу
	if reflect.TypeOf(number) == reflect.TypeOf(c) {
		fmt.Println(" Type OK")
	} else {
		fmt.Println(" Type Not OK")
	}

}
