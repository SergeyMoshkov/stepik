package main

import "fmt"

func lengthArray () {
  var a, c, len int
  
  for a >= 0 {
    fmt.Scan(&a)
    if a == 0 {
      break
    } else if a > c {
        len = 1
        c = a
    } else {
      if a == c {
        len += 1
      }
    }
    fmt.Println(len)
  }
}

func main() {
  lengthArray()
}
