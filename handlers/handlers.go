package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error")
	}
}
