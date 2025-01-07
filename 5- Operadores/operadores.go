package main

import "fmt"

func main() {
	// ARITMETICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// FIM DOS ARITMETICOS

	// ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println("---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("---------------------")
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNÁRIOS

	fmt.Println("---------------------")
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}