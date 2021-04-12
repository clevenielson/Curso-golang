package main

import "fmt"

func generica(interf interface{}) {
  fmt.Println(interf)
}

func main() {
  generica("String")
  generica(1)
  generica(true)

  fmt.Println(1, 2, "string", false, true, float64(12345))

  mapa := map[interface{}] interface{}{
    1:            "string",
    float64(100): true,
    "string":     "string",
    true:         float64(12),
  }
  fmt.Println(mapa)
}
