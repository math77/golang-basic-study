package main

import (
  "errors"
  "fmt"
)

func main(){
  pilha := Pilha{}

  fmt.Println("Nova pilha criada com o tamanho de: ", pilha.Tamanho())
  fmt.Println("Pilha vazia? ", pilha.Vazia())

  pilha.Empilhar("Golang")
  pilha.Empilhar("2012")
  pilha.Empilhar(12.56)
  pilha.Empilhar(11)
  pilha.Empilhar("Led Zeppelin")
  pilha.Empilhar(-90)

  fmt.Println("Tamanho da pilha agora: ", pilha.Tamanho())
  fmt.Println("Pilha vazia? ", pilha.Vazia())

  // Enquanto a pilha não estiver vazia
  for !pilha.Vazia() {
    v, _ := pilha.Desempilhar()
    fmt.Println("Desempilhando ", v)
    fmt.Println("Tamanho: ", pilha.Tamanho())
    fmt.Println("Vazia? ", pilha.Vazia())
  }

  _, err := pilha.Desempilhar()
  if err != nil {
    fmt.Println("Erroor: ", err)
  }
}


type Pilha struct {
  valores []interface{}
}

func (pilha Pilha) Tamanho() int {
  return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
  return pilha.Tamanho() == 0
}

// *Pilha indica um ponteiro para um objeto do tipo Pilha.
func (pilha *Pilha) Empilhar(valor interface{}) {
  pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
  if pilha.Vazia() {
    return nil, errors.New("Pilha vaziaaa")
  }

  valor := pilha.valores[pilha.Tamanho()-1]
  pilha.valores = pilha.valores[:pilha.Tamanho()-1]
  return valor, nil
}
