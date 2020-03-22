package main

func main() {

	// Estrutura específica para trabalhar com concorrência
	select {
	case t1 := <-c1:
		return t1
	}

}
