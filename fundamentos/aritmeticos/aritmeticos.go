package main

import "fmt"

func main() {
	a := 3
	b := 2
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b) 
	fmt.Println("Divisão =>", a/b) 
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b) 

	// bilwise 
	fmt.Println("E =>", a&b) // 11 & 10=10
	fmt.Println("Ou =>", a |b) // 11| 10 = 11
}