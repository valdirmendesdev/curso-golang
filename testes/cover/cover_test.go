package cover

import "testing"

func TestSomar(t *testing.T) {
	if somar(2, 2) != 4 {
		t.Error("Erro na operação de somar")
	}
}

//go test -cover -> Apresenta a cobertura de testes no terminal
//go test -coverprofile=resultado.out -> Gera um arquivo com a cobertura
//go tool cover -func=resultado.out -> Apresenta o resultado de cobertura a partir de um arquivo
//go tool cover -html=resultado.out -> Gera um HTML com o relatório de cobertura
