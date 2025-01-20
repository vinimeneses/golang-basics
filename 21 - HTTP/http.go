package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Olá Mundo"))
	if err != nil {
		log.Fatal(err)
	}
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Usuários do Sistema: x..y..z.."))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
