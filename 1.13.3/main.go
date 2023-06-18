package main

import (
	"fmt"
	"time"
)

func main() {
	var k int
	fmt.Print("input k: ")
	fmt.Scan(&k)

	h := k / 3600      
	m := (k % 3600) / 60 

	fmt.Printf("It is %d hours %d minutes.\n", h, m)

	// Вывод текущего времени
	now := time.Now()
	fmt.Println("Текущее время:", now.Format("15:04:05"))
}
