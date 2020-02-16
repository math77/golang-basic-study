package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
  l := *lista
  removido := l[indice]
  *lista = append(l[0:indice], l[indice+1:]...)
  return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
  return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
  return lista.RemoverIndice(len(*lista)-1)
}

func main(){
    lista := ListaGenerica{1, "Café", 2.3, true, "Matheus", false, 77}

    fmt.Printf("Lista original:\n%v\n\n", lista)
    fmt.Printf("Removendo do inicio: %v, ficando:\n%v\n ", lista.RemoverInicio(), lista)
}
