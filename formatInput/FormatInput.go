package main

import "fmt"

// FormatInput function formatted output function
func FormatInput() {
	var number float64
	_, err := fmt.Scan(&number)
    if err != nil {
        fmt.Println(err)
    }

	// execution using the switch construct
	switch {
	case number <= 0:
		fmt.Printf("число %.2f не подходит\n", number)
	case number > 10000:
		fmt.Printf("%e\n", number)
	default:
		fmt.Printf("%.4f\n", number*number)
	}

	// execution using the if construct
	if number <= 0 {
        fmt.Printf("число %.2f не подходит\n", number)
	} else if number > 10000 {
        fmt.Printf("%e\n", number)
	} else {
        fmt.Printf("%.4f\n", number*number)
	}
}

func main()  {
	FormatInput()
}
