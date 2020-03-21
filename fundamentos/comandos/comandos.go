package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!", "Go")
}

/*
go help [command] => retorna a documentação de um comando

godoc -http=:6060 => Disponibiliza a documentação da linguagem offline localhost:6060

go doc cmd/vet => Documentação do comando vet

go vet [filename] => verifica o código daquele arquivo

go get -u [url] => instala um pacote utilizando o GOPATH para armazenar
*/
