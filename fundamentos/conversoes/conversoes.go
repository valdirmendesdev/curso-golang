package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//Cuidado...
	//Não está convertendo para string, está imprimindo o caracter da tabela unicode
	fmt.Println("teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	//conversão de boolean
	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("verdadeiro")
	}
}
