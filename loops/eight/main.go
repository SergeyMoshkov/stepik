package main

import (
  "fmt"
  "strconv"
)

func main() {
  var x, y string
  var a, b int
  fmt.Scan(&x, &y)
  for _, v := range x {
    xstr:=string(v)
    a,_=strconv.Atoi(xstr)
    for _, v := range y {
      xstr:=string(v)
      b,_=strconv.Atoi(xstr)
      if a == b {
        fmt.Print(a, " ")
      }
    }
  }
}
