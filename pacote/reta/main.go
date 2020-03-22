package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	fmt.Printf("%.2f", Distancia(p1, p2))
}
