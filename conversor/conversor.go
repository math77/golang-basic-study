package main


import (
  "fmt"
  "os"
  "strconv"
)

func main(){

  if len(os.Args) < 3 {
    fmt.Println("Uso: conversor <valores> <unidades>")
    os.Exit(1)
  }

  var unidadeDestino string
  unidadeOrigem := os.Args[len(os.Args)-1]
  valoresOrigem := os.Args[1 : len(os.Args)-1]

  if unidadeOrigem == "celsius" {
    unidadeDestino = "fahrenheit"
  } else if unidadeOrigem == "quilometros" {
    unidadeDestino = "milhas"
  } else {
    fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
    os.Exit(1)
  }

  //range aplicado a slice retorna um valor para indice e outro para o valor
  for i, v := range valoresOrigem {

    valorOrigem, err := strconv.ParseFloat(v, 64)

    if err != nil {
      fmt.Printf("O valor %s na posição %d não é um número válido\n", v, i)
      os.Exit(1)
    }

    var valorDestino float64

    if unidadeOrigem == "celsius" {
      valorDestino = valorOrigem*1.8 + 32
    } else {
      valorDestino = valorOrigem / 1.60934
    }

    fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
  }
}
