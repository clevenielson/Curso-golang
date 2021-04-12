package main

import "fmt"

type pessoa struct {
  nome      string
  sobrenome string
  idade     uint8
  altura    uint8

}

type estudante struct {
  pessoa
  Curso     string
  Faculdade string
}


func main() {
  fmt.Println("Herança")

  p1 := pessoa{"Clevin", "Ben", 20, 178}
  fmt.Println(p1)

  e1 := estudante{p1, "Engenharia", "USP"}
  fmt.Println(e1.altura)
}
