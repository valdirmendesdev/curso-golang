package main

import (
	"fmt"
	"math"
)

func operacao(operacao string, result interface{}) {
	fmt.Printf("%s => %v\n", operacao, result)
}

func main() {
	a := 3
	b := 2

	operacao("Soma", a+b)
	operacao("Subtração", a-b)
	operacao("Divisão", a/b)
	operacao("Multiplicação", a*b)
	operacao("Módulo", a%b)

	fmt.Println()

	//bitwise
	operacao("E", a&b)   // 11 & 10 = 10
	operacao("Ou", a|b)  // 11 | 10 = 11
	operacao("Xor", a^b) // 11 ^ 10 = 01

	fmt.Println()

	//outras operações usando math...
	c := 3.0
	d := 2.0
	operacao("Maior", math.Max(float64(a), float64(b)))
	operacao("Menor", math.Min(c, d))
	operacao("Exponenciação", math.Pow(c, d))

}
