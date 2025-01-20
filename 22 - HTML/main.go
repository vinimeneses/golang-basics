package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := Usuario{"Fulano", "fulano@gmail.com"}
		err := templates.ExecuteTemplate(w, "home.html", u)
		if err != nil {
			return
		}
	})

	fmt.Println("Server is running at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
