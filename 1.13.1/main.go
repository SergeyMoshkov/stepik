package main

import (
	"fmt"
)

func main() {
	var i int
	var s int
    fmt.Scan(&i)
	a := i / 100
	s = a % 10
	b := i / 10
	c := b % 10
	s = s + c
	d := i % 10
	s = s + d
}