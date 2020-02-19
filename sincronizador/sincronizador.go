package main

import (
  "fmt"
  "time"
  "math/rand"
  "sync"
)

func main(){
  inicio := time.Now()
  rand.Seed(inicio.UnixNano())

  //Var utilizada para sincronizar todas as goroutines.
  var controle sync.WaitGroup

  for i := 0; i < 5; i++ {
    //Indica que uma nova goroutine deve ser sincronizada.
    controle.Add(1)
    //Passamos a var controle como um ponteiro pois precisa ser a mesma em todas as goroutines.
    go executar(&controle)
  }

  // Travar execução do main() até que todas as goroutines tenham acabado.
  controle.Wait()

  fmt.Printf("Finalizando em %s.\n", time.Since(inicio))
}

func executar(controle *sync.WaitGroup) {
  defer controle.Done()

  duracao := time.Duration(1+rand.Intn(5)) * time.Second
  fmt.Printf("Dormindo por %s...\n", duracao)
  time.Sleep(duracao)

}
