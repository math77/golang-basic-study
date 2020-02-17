package main

import (
  "fmt"
  "time"
)

func GerarFibonacci(n int) func() {
  return func() {
    a, b := 0, 1

    fib := func() int {
      a, b = b, a+b

      return a
    }

    for i := 0; i < n; i++ {
      fmt.Printf("%d ", fib())
    }
  }
}

func Cronometrar(funcao func()){
  //armazenar o tempo atual.
  inicio := time.Now()
  funcao()
  //Since -> calcula intervalo entre momento atual e o momento armazenado em inicio.
  //Retorna um valor do tipo time.Duration
  fmt.Printf("\nTempo de execução: %s\n\n", time.Since(inicio))
}

func main() {
  Cronometrar(GerarFibonacci(8))
  Cronometrar(GerarFibonacci(48))
  Cronometrar(GerarFibonacci(88))
  Cronometrar(GerarFibonacci(108))
}
