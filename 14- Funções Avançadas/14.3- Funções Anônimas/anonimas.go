package main

import "fmt"

func main() {

	func() {
		println("Olá Mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")
	
	fmt.Println(retorno)
}