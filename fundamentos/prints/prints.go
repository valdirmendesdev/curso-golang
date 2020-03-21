package main

import "fmt"

func main() {
	//Print não quebra linha.
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//Println - continua na mesma linha e ao final quebra linha
	fmt.Println(" Nova")
	fmt.Println(" linha")

	x := 3.141516

	//Não vai deixar concatenar
	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs) //Uma solução
	fmt.Println("O valor de x é", x)    //Outra solução

	fmt.Printf("O valor de x é %.2f", x) //Formatando um float com duas casas decimais

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
