package main

/*
Algorithm
1. Check if p**n divides b
2. Solve bx = 1 (mod p)
3. If x_n is a solution then set a_n = c * x_n (mod p)
4. Set a_n b = r(N) - r(N+1) * p (mod p)
5. Check a_n * b = r(N) (mod p)
6. A_n = c*r_n (mod p)
7.if r(N+1) = r_i for any I in N then {a_0, a_1, ..., a_n} you are done
8. If not 7, set lambda = n-i +1 and go back to 2
*/

func padic_expansion(a, b, p int64) *Padic {
  var n int64 = 1
  g := GCD(b, power(p,n))
  for g != 1 {
    g = GCD(b, power(p, n+1))
    n = n+1
  }
  b %= p
  
  iter := make([]int64, 10)
  var y int64
  for y = a; y <= n+a; y++ {
    if b*iter[y-a]-1 == p*n {
      iter = append(iter, a*y)
    } else {
      iter = append(iter, a*y-p*n)
    }
    if iter[n-1]-a*y == 0 && y != 0 {
      a = a - iter[n-1]*b
    } else {
      a /= p
    }
    n += 1
  }
  for a < 0 || a > p-1 {
    if a < 0 {
      a = a +p
    } else if a > p-1 {
      a = a-p
    }
  }
  return &Padic{
    num: [2]int64{a,b}
    norm: valuation(p,n),
    seq: padic_code(a, p, iter),
  }
}
