package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (receptor)

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 15000.00, 0.1}
	produto2.preco = 200
	fmt.Println(produto2, produto2.precoComDesconto())

}
