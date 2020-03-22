package main

import "math"

//Iniciando o nome de qualquer objeto com letra maiúscula torna-o Público
//Vísivvel dentro e fora do pacote
//Iniciando o nome de qualquer objeto com letra minúsculo torna-o Privado,
//dentro do pacote

//Exemplos
// Ponto -> Será público
// ponto ou _Ponto -> Será privado

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
