package main

import "fmt"

func inverterSinal(numero int) int {
  return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
  *numero = *numero * -1
}

func main() {
  numero := 20
  inverteNumero := inverterSinal(numero)
  fmt.Println(inverteNumero)
  fmt.Println(numero)

  novoNumero := 40
  fmt.Println(novoNumero)
  inverterSinalComPonteiro(&novoNumero)
  fmt.Println(novoNumero)
}
