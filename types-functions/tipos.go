package main

import "fmt"

/*
  Function type são semelhantes a interfaces.
  Qualquer função que receba dois inteiros e retorne outro inteiro pode usar
  o tipo Agregadora.
*/
type Agregadora func(n, m int) int

func Agregar(valores []int, valorInicial int, fn Agregadora) int {
  agregado := valorInicial

  for _, v := range valores {
    agregado = fn(v, agregado)
  }

  return agregado
}

func CalcularSoma(valores []int) int {
  soma := func(n, m int) int {
    return n+m
  }

  return Agregar(valores, 0, soma)
}

func CalcularProduto(valores []int) int {
  produto := func(n, m int) int {
    return n*m
  }

  return Agregar(valores, 1, produto)
}

func main() {
  valores := []int{3, 2, 5, 7, -3, 22, -1}
  fmt.Println(CalcularSoma(valores))
  fmt.Println(CalcularProduto(valores))
}
