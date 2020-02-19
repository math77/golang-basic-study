package main

import (
  "fmt"
  "time"
)

/*
  Podemos utilizar a funcao After() para controlar o tempo de execução de uma
  goroutine. Com isso podemos tomar medidas caso alguma goroutine não execute no tempo
  desejado.
*/

func executar(c chan<- bool) {
  time.Sleep(1 * time.Second)
  c <- true
}

func main() {
  c := make(chan bool, 1)
  go executar(c)
  fmt.Println("Esperando...")


  fim := false

  for !fim {
    select {
    case fim = <-c:
      fmt.Println("Fim!")
    case <-time.After(2 * time.Second):
      fmt.Println("Timeout!")
      fim = true
    }
  }
}
