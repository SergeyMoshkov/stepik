package main

import "fmt"

func main() {
    var N int
    var m int

    fmt.Scan(&N)
    arr := make([]int, N, N)

    for i := 0; i < 5; i++ {
        fmt.Scan(&m)
        arr[i] = m
    }
    fmt.Println(arr[3])
}