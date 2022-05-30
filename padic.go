package main

type Padic struct {
  num [2]int64
  val int64
  norm [2]int64
  seq []int64
}

func valuation(prime, val int64) [2]int64{
  norm := [2]int64{1, power(prime, val)}
  return norm
}

func power(base, exponent int64) int64 {
  var output int64 = 1
  for exponent != 0 {
    output *= base
    exponent -= 1
  }
  return output 
}

func GCD(a, b int64) int64 {
  var gcd int64
  var i int64
  for i = 1; i <= a && i <= b; i++ {
    for a%i == 0 && b%i == 0 {
      gcd = i
    }
  }
  return gcd
}

func padic_code(a, p int64, iter []int64) []int64 {
  u := make([]int64, 4)
  for i :=0; i < len(iter); i++ {
    a += iter[i]
    if i >0 && iter[i] != iter[i-1] {
      if iter[i] != 0 {
        u = append(u, iter[i]
      }
    }
  }
  for i := 0; i < len(u); i++ {
     x := u[i]
     for x < 0 || x > p-1 {
       if x < 0 {
         x = x+p
       } else if x >p-1 {
         x = x-p
       } else if x== 0 {
         continue
       }
   u[i] = x
  }
  return u
  }
                   
               
