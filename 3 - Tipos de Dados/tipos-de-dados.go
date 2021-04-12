package main

import (
  "fmt"
  "errors"
)

func main() {
  // int, int8, int16, int32, int64

  numero := 10249
  fmt.Println(numero)
  fmt.Printf("%T\n", numero)

  // unsygned int (int sem sinal)
  // uint, unit8, unit16, uint32 uint64

  var numero2 uint16 = 10000
  fmt.Println(numero2)
  fmt.Printf("%T\n", numero2)

  // alias - Apelido
  // INT32 = RUNE

  var numero3 rune = 12456
  fmt.Println(numero3)
  fmt.Printf("%T\n", numero3)

  // BYTE = UINT8
  var numero4 byte = 123
  fmt.Println(numero4)
  fmt.Printf("%T\n", numero4)

  var numeroReal1 float32 = 123.45
  fmt.Println(numeroReal1)
  fmt.Printf("%T\n", numeroReal1)

  var numeroReal2 float64 = 123000.45
  fmt.Println(numeroReal2)
  fmt.Printf("%T\n", numeroReal2)

  numeroReal3 := 12345.67
  fmt.Println(numeroReal3)
  fmt.Printf("%T\n", numeroReal3)

  // Fim Numeros Reais

  var str string = "Texto"
  fmt.Println(str)

  str2 := "Texto2"
  fmt.Println(str2)

  char := 'H'
  fmt.Println(char)

  // FIM String

  var texto uint16
  fmt.Println(texto)

  var booleano1 bool = true
  fmt.Println(booleano1)

  var booleano2 bool = false
  fmt.Println(booleano2)

  var erro error = errors.New("Erro interno")
  fmt.Println(erro)
}
