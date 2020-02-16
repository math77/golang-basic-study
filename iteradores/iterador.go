package main

import "fmt"

func main() {

  a, b := 0, 10

  //uso semelhante ao while
  for a < b {
    a += 1
    fmt.Println(a)
  }

  for i := 0; i < 10; i++ {
    fmt.Println("Olá")
  }

  l := 10
  for ; l < 20; l++ {
    fmt.Println(l)
  }

  numeros := []int{1, 2, 4, 6, 10, 20}
  //Quando precisamos apenas do indice dos valores.
  for indice := range numeros {
    numeros[indice] *= 2
  }

  fmt.Println(numeros)

}
