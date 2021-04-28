package controllers

import (
  "api/src/banco"
  "api/src/modelos"
  "api/src/repositorios"
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "fmt"
)

// CriarUsuario insere um usuario no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
  corpoRequest, erro := ioutil.ReadAll(r.Body)
  if erro != nil {
    log.Fatal(erro)
  }

  var usuario modelos.Usuario
  if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
    log.Fatal(erro)
  }

  db, erro := banco.Conectar()
  if erro != nil {
    log.Fatal(erro)
  }

  repositorio := repositorios.NovoRepositorioDeUsuarios(db)
  usuarioID, erro := repositorio.Criar(usuario)
  if erro != nil {
    log.Fatal(erro)
  }
  
  w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}

// BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Buscando todos os usuário!"))
}

// BuscarUsuario busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Buscando um usuário!"))
}

// AtualizarUsuario altera as informacoes de um usuario no DB
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario exclui as informacoes de usuario no BD
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Deletando Usuário!"))
}
