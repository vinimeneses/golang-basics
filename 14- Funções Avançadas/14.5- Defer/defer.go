package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println("Fim do programa")
}