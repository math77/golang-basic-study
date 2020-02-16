package main

import "fmt"

func main(){

  fib := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 23, 12, 20}

  fmt.Println(fib)
  fmt.Println(fib[:3])
  fmt.Println(fib[2:])
  fmt.Println(fib[:])

  numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

  //inserir no final do slice
  numeros = append(numeros, 110)
  fmt.Println(numeros)

  //inserir no inicio do slice
  novosNumeros := []int{0, 5}
  numeros = append(novosNumeros, numeros...)
  fmt.Println(numeros)

  //inserir em qualquer parte do slice
  numerosMeio := []int{55, 56, 57}
  numeros = append(numeros[:6], append(numerosMeio, numeros[6:]...)...)
  fmt.Println(numeros)
}
