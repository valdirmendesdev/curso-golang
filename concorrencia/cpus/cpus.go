package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Número de processadores da minha máquina:", runtime.NumCPU())
}
