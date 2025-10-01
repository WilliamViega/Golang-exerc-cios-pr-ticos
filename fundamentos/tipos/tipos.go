package main

import (
    "fmt"
    "math"
    "reflect"
)

func main() {
    // números inteiros
    fmt.Println(1, 2, 1000)
    fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

    // sem sinal (só positivos)... uint8 uint16 uint32 uint64
    var b byte = 255
    fmt.Println("O byte é", reflect.TypeOf(b))

    // com sinal... int8 int16 int32 int64
    il := math.MaxFloat64
    fmt.Println("O valor máximo do float64 é", il)

    var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
    fmt.Println("O rune é", reflect.TypeOf(i2))

    // string com múltiplas linhas
    s2 := `Olá 
meu 
nome
é 
Leo`
    fmt.Println(s2)
}
 