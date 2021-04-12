package main

import (
  "fmt"
  "time"
)

func main() {
  // Concorrencia != Paralelismo
  go escrever("Olá mundo!") //goroutine
  escrever("programando em Go!")
}

func escrever(texto string) {
  for {
    fmt.Println(texto)
    time.Sleep(time.Second)
  }
}
