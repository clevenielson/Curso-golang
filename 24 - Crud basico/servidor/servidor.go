package servidor

import (
  "crud/banco"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

type usuario struct {
  ID    uint32 `json:"id"`
  Nome  string `json:"nome"`
  Email string `json:"email"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
  corpoRequisicao, erro := ioutil.ReadAll(r.Body)
  if erro != nil {
    w.Write([]byte("Falha ao ler o corpo da requisicao!"))
    return
  }

  var usuario usuario
  if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
    w.Write([]byte("Erro ao converter o usuario para struct"))
    return
  }

  db, erro := banco.Conectar()
  if erro != nil {
    w.Write([]byte("Erro ao conectar no banco de dados"))
    return
  }
  defer db.Close()

  statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
  if erro != nil {
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }
  defer statement.Close()

  insercao, erro := statement.Exec(usuario.Name, usuario.Email)
  if erro != nil {
    w.Write([]byte("Erro ao executar o statement!"))
    return
  }

  idInserido, erro := insercao.LastInsertID()
  if erro != nil {
    w.Write([]byte("Erro ao obter o id inserido!"))
    return
  }

  w.WriteHeader(http.StatusCreated)
  w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! id: %d", idInserido)))

}
