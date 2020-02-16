package main

import (
  "fmt"
  "reflect"
)

type ListaDeCompras []string

func imprimirSlice(slice []string){
  fmt.Println("Slice: ", slice)
  fmt.Println(reflect.TypeOf(slice).String())
}

func imprimirLista(lista ListaDeCompras) {
  fmt.Println("Lista de compras: ", lista)
  fmt.Println(reflect.TypeOf(lista).String())
}

func main() {
  lista := ListaDeCompras{"Alface", "Atum", "Tomate"}
  slice := []string{"Alface", "Atum", "Tomate"}

  imprimirSlice([]string(lista))
  imprimirLista(ListaDeCompras(slice))
}
