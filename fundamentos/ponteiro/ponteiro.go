package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiros
	var p *int = nil
	p = &i
	*p++
	fmt.Println(*p)
	fmt.Println(p, *p, i, &i)

}
