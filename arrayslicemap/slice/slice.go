package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array! Slice define um pedaço de um array
	//Slice não cria nova referência, aponta para memória do array
	s2 := a2[1:3] // slice que contém valores da posição até 3 (não incluindo a 3)
	fmt.Println(a2, s2)

	s3 := a2[:4] // slice que contém os primeiros 4 primeiros valores do array
	fmt.Println(a2, s3)

	//slice aponta para o tamanho e ponteiro para elementos de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)

	// Prova de que o slice está apontando para a mesma referência de memória do array
	a2[1] = 20
	fmt.Println(s2, s3, s4)

}
