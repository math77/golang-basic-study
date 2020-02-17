package main

import (
  "fmt"
  "regexp"
  "strings"
)

func main() {
  expr := regexp.MustCompile("\\b\\w")

  transformadora := func(s string) string {
    return strings.ToUpper(s)
  }

  texto := "antonio carlos jobim"

  //utiliza a função guardada em transformadora
  fmt.Println(transformadora(texto))
  //utiliza a função com a expressão guardade em expr.
  fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}
