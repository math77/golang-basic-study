package main

import (
  "fmt"
  "net/http"
  "time"
)

func main(){
  //registra nosso serviço com a rota que ele irá atender e definindo uma função
  //que atenda a assinatura do http.HandlerFunc
  http.HandleFunc("/tempo", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
  })
  //especifica que o servidor deve aceitar reqs. na porta 8080.
  http.ListenAndServe(":8080", nil)
}
