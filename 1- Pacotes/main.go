package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrevendo()

	erro := checkmail.ValidateFormat("viniciusmenesesdev@gmail.com")
	fmt.Println(erro)
}
