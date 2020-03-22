package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo, primeiro = p2, p1
	return //Retorno limpo, uma vez que nomeei os parametros de retorno
}

func main() {
	fmt.Println(trocar(2, 1))
}
