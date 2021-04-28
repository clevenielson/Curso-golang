package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
  "fmt"
  "strconv"
)

var (
  // StringConexaoBanco é a string de conexão com o Mysql
  StringConexaoBanco = ""

  // Porta onde a API vai estar rodando
  Porta = 0
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
  var erro error

  if erro = godotenv.Load(); erro != nil {
    log.Fatal(erro)
  }

  Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
  if erro != nil {
    Porta = 9000
  }

  StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
    os.Getenv("DB_USUARIO"),
    os.Getenv("DB_SENHA"),
    os.Getenv("DB_NOME"),
  )
}
