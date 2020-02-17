package main

import "fmt"

func main() {
  a, b := 0, 1

  //pode utilizar o contexto da função main porque foi definida dentro dela.
  fib := func() int {
    a, b = b, a+b
    return a
  }

  for i := 0; i < 8; i++ {
    fmt.Printf("%d \n", fib())
  }
}
