package main

import "fmt"

func main() {
	x, y := 1, 2

	//apenas postfix
	x++ //x += 1 ou x = x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

}
