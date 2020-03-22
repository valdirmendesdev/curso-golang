package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.98,
			"Guga Pereira":   9876.09,
		},
		"J": {
			"José João": 654987.23,
		},
	}
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
