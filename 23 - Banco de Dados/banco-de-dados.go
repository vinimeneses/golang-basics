package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	erro = db.Ping()
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	lines, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(lines)
	defer lines.Close()

}
