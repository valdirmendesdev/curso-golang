package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":     11235.45,
		"Gabriel Silva": 15456.78,
	}

	funcsESalarios["Novo funcionário"] = 10000.01

	delete(funcsESalarios, "Não existe")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
