package main

// CRUD - CREATE,  READ, UPDATE, DELETE

// Create - Post
// Read - Get
// Update - Put
// Delete - Delete

import (
  "crud/servidor"
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {

  router := mux.NewRouter()
  router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

  fmt.Println("Escutando na porta 5000")
  log.Fatal(http.ListenAndServe(":5000", router))

}
