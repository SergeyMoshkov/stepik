package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Введите начало отрезка a: ")
	fmt.Scan(&a)
	fmt.Print("Введите конец отрезка b: ")
	fmt.Scan(&b)

	max := -10000000  // Инициализируем переменную max с отрицательным значением, чтобы проверить, было ли найдено кратное 7 число

	// Проверяем, если a больше b, меняем их значения местами
	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			if i > max {
				max = i
				
			}
		}
	}

	if max == 0 {
		fmt.Println("NO")
	} else {
		fmt.Printf("Самое большое число, кратное 7, на отрезке от %d до %d: %d\n", a, b, max)
	}
}
