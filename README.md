# Estudos iniciais na linguagem Go

Utilizando o livro da Casa do Código para entender os principios básicos
da linguagem.
[Link para comprar o livro] (https://www.casadocodigo.com.br/products/livro-google-go?_pos=1&_sid=a5a6fc573&_ss=r)

### Conceitos básicos

Os arquivos de código em Go seguem a seguinte estrutura:

1. Declaração do pacote
   1. ```go
        package main
      ```
2. Declaração dos pacotes externos que o arquivo atual depende
   2. ```go
        import (
          "fmt"
          "os"
          )
      ```
3. Código referente ao que está sendo escrito.
   3. ```go
        func main(){}
      ```
