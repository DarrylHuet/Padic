package main

import (
  "fmt"
)

//Currently 22/7 computes to 3, since 22/7 = 3.14285714286
//Under the Padic, our fraction should be an exact computation with specific precision
func main() {
  var p int64 = 13
  var a int64 = 22
  var b int64 = 7
  
  padic := padic_expansion(a,b,p)
  fmt.Println("Padic Rational for 22/7 mod 13: ", *padic)
}
