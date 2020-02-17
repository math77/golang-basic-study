package main

import (
  "fmt"
  "os"
)

func CriarArquivos(dirBase string, arquivos ...string) {
  for _, nome := range arquivos {
    caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")

    //Retorna um "descritor" do tipo os.File que pode ser usado para outras tarefas.
    //Pode retornar um erro caso não seja possivel criar o arquivo.
    arq, err := os.Create(caminhoArquivo)

    //defer -> Realizar tarefa antes da atual retornar. Atrasar a execução de algo.
    defer arq.Close()

    if err != nil {
      fmt.Printf("Erro ao criar arquivo %s: %v\n", nome, err)
      os.Exit(1)
    }

    fmt.Printf("Arquivo %s criado.\n", arq.Name())
  }
}

func main() {
  //retorna o diretorio atual.
  current, err := os.Getwd()

  if err != nil {
    os.Exit(1)
  }

  CriarArquivos(current)
  CriarArquivos(current, "teste")
  CriarArquivos(current, "teste1", "teste2")


}
