package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)
	go escrever("Eu vou ser um dev GoLang!", canal)
	
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second * 1)
	}

	close(canal)
}
