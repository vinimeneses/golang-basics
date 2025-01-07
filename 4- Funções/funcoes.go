package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	resultado := somar(10, 20)
	println(resultado)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Texto da função")
	println(result)
}