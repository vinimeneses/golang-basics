package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Rua dos Bobos")
	fmt.Println(TipoEndereco)
}
