package main

import "fmt"

func main() {
	s := make([]int, 10) //Cria um slice e um array interno com 10 posições
	s[9] = 12
	fmt.Println(s)

	//Cria um slice com 10 posições e um array interno com 20 posições
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) //cap => retorna a capacidade

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	//Quando o array já está em sua capacidade máxima e adicionamos um elemento
	//a capacidade dobra de tamanho
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
