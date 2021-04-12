package main

import (
  "fmt"
  "time"
  "sync"
)

func main() {
  var waitGroup sync.WaitGroup

  waitGroup.Add(4)

  go func() {
    escrever("Olá Mundo!")
    waitGroup.Done() // -1
  }()

  go func() {
    escrever("Programando em Go!")
    waitGroup.Done() // -1
  }()

  go func() {
    escrever("Goroutine 3")
    waitGroup.Done() // -1
  }()

  go func() {
    escrever("Goroutine 4!")
    waitGroup.Done() // -1
  }()

  waitGroup.Wait() // 2

  escrever("Olá Mundo!")
  escrever("Programando em Go!")
}

func escrever(texto string) {
  for i := 0; i < 3; i++ {
    fmt.Println(texto)
    time.Sleep(time.Second)
  }
}
