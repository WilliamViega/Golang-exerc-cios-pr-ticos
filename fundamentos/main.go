package main

import "fmt"

func somar(a int, b int) int {
    return a + b
}

func imprimir(valor int) {
    fmt.Println("Resultado:", valor)
}

func main() {
    resultado := somar(3, 4)
    imprimir(resultado)
} 