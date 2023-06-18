package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := (n % 10)

	if 11 <= n && n <= 20 {
		fmt.Println(n, "korov")
	} else if a == 1 {
		fmt.Println(n, "korova")
	} else if a == 2 || a == 3 || a == 4 {
		fmt.Println(n, "korovy")
	} else {
		fmt.Println(n, "korov")
	}
}
