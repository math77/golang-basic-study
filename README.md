# Estudos iniciais na linguagem Go

Utilizando o livro da Casa do Código para entender os principios básicos
da linguagem.
[Link para comprar o livro!](https://www.casadocodigo.com.br/products/livro-google-go?_pos=1&_sid=a5a6fc573&_ss=r)

### Conceitos básicos

Os arquivos de código em Go seguem a seguinte estrutura:

1. Declaração do pacote
2. Declaração dos pacotes externos que o arquivo atual depende
3. Código referente ao que está sendo escrito.

* Todo código em Go deve obrigatoriamente existir em um package
* Todo programa em Go deve ter um pacote principal com uma função main() responsável por iniciar o programa.

#### Declarar variáveis

Para declarar váriaveis em Go podemos fazer de duas formas:

Seguindo a estrutura: `var nome tipo`

```go
  var nomeVariavel string
```
Ou, declarando e já atribuindo um valor:

```go
  nome := "Matheus"
```

**Obs.: Variáveis declaradas e não utilizadas geram erro de compilação.**
