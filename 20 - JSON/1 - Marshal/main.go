package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Zoe", "Golden Retriever", 2}

	cachorroEmJSON, erro := json.MarshalIndent(c, "", " ")
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Nick",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.MarshalIndent(c2, "", " ")
	cachorro3EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorro3EmJSON))
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
