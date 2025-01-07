package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)
	fmt.Println("----")
	escrever("Olá Mundo", 1, 2, 3, 4, 5, 6, 7)
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		println(texto, numero)
	}

}