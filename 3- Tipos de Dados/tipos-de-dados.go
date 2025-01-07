package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = rune

	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := 1 > 2
	fmt.Println(booleano2)
	
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}