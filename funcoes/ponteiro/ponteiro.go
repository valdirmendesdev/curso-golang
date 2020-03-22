package main

import "fmt"

// Ponteiro é representado por um *

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n) //Passando parâmetro por valor
	fmt.Println(n)

	//& usado para obter o endereço da variável
	inc2(&n)
	fmt.Println(n)
}
