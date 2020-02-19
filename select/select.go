package main

import "fmt"

func separar(nums []int, i, p chan<- int, pronto chan<- bool) {
  for _, n := range nums {
    //separa os numeros e enviar para os respectivos canais.
    if n%2 == 0 {
      p <- n
    } else {
      i <- n
    }
  }
  //indica que o processamento acabou enviando um valor para o canal.
  pronto <- true
}

func main(){
  i, p := make(chan int), make(chan int)
  pronto := make(chan bool)
  nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

  go separar(nums, i, p, pronto)

  var impares, pares []int
  fim := false

  for !fim {
    select {
    case n := <-i:
      impares = append(impares, n)
    case n := <-p:
      pares = append(pares, n)
    case fim = <-pronto:
    }
  }

  fmt.Printf("Impares: %v | Pares: %v\n", impares, pares)
}
