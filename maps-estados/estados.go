package main

import "fmt"

func main() {

  estados := make(map[string]Estado, 4)
  estados["RN"] = Estado{"Rio Grande Do Norte", 3373960, "Natal"}
  estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
  estados["PB"] = Estado{"Paraiba", 3914418, "João Pessoa"}
  estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

  fmt.Println(estados)

  delete(estados, "GO")
  fmt.Println(estados)

  for sigla, estado := range estados {
    fmt.Printf("%s (%s) possui %d habitantes\n", estado.nome, sigla, estado.populacao)
  }
}


type Estado struct {
  nome string
  populacao int
  capital string
}
