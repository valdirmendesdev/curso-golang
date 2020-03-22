package main

import "fmt"

func main() {
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123456789] = "name1"
	aprovados[987654321] = "name2"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])
	delete(aprovados, 123456789) //exclui uma chave de um map
	fmt.Println(aprovados)
	fmt.Println(len(aprovados))
}
