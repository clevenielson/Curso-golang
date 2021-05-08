package controllers

import (
  "net/http"
  "io/ioutil"
  "api/src/respostas"
  "api/src/modelos"
  "encoding/json"
  "api/src/banco"
  "api/src/seguranca"
  "api/src/repositorios"
)

// Login é responsavel por autenticar usuario na api
func Login(w http.ResponseWriter, r *http.Request) {
  corpoRequisição, erro := ioutil.ReadAll(r.Body)
  if erro != nil {
    respostas.Erro(w, http.StatusUnprocessableEntity, erro)
    return
  }

  var usuario modelos.Usuario
  if erro = json.Unmarshal(corpoRequisição, &usuario); erro != nil {
    respostas.Erro(w, http.StatusBadRequest, erro)
    return
  }

  db, erro := banco.Conectar()
  if erro != nil {
    respostas.Erro(w, http.StatusInternalServerError, erro)
    return
  }
  defer db.Close()

  repositorio := repositorios.NovoRepositorioDeUsuarios(db)
  usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
  if erro != nil {
    respostas.Erro(w, http.StatusInternalServerError, erro)
    return
  }

  if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
    respostas.Erro(w, http.StatusUnauthorized, erro)
    return
  }

  w.Write([]byte("Você está logado! Parabéns!"))

}
