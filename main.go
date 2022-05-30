package main

import (
  "fmt"
)

func main() {
  var p int64 = 13
  var a int64 = 22
  var b int64 = 7
  
  padic := padic_expansion(a,b,p)
  fmt.Println("Padic Rational for 22/7 mod 13: ", *padic)
}
