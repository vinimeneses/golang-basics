package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10

	for i > 0 {
		println(i)
		i--
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

    for index, nome := range nomes {
        fmt.Printf("Índice: %d, Nome: %s\n", index, nome)
    }

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}
