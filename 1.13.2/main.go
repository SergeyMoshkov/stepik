package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var s int
	var t int
	var v int
    fmt.Scan(&i)
	s = i % 10
	t = (i / 10) % 10
	v = (i / 100) % 10
	fmt.Printf(strconv.Itoa(s) + strconv.Itoa(t) + strconv.Itoa(v))
}
