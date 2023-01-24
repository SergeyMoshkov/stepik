package main

import "fmt"

func main() {
  var x, p, y, year float32
  fmt.Scan(&x, &p, &y)
  for year = 0; x < y; year += 1 {
    var result int
    x = x + ((x / 100) * p)
    result = int(x)
    x = float32(result)
  }
  fmt.Println(year)
}
