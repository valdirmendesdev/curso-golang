package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	//Laço infinito
	for {
		fmt.Println("Executando eternamente...")
		time.Sleep(time.Second)
	}

}
