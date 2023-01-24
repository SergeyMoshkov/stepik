package main

import "fmt"

func FormatInput() {
	var number float64
	_, err := fmt.Scan(&number)
    if err != nil {
        fmt.Println(err)
    }
//	switch {
//	case number <= 0:
//		fmt.Printf("число %.2f не подходит\n", number)
//	case number > 10000:
//		fmt.Printf("%e\n", number)
//	default:
//		fmt.Printf("%.4f\n", number*number)
//	}

	if number <= 0 {
        fmt.Printf("число %.2f не подходит\n", number)
	} else if number > 10000 {
        fmt.Printf("%e\n", number)
	} else {
        fmt.Printf("%.4f\n", number*number)
	}
}
