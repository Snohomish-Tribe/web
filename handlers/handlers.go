package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	http.FileServer(http.Dir("static"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error")
	}
}
func About(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/about.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error")
	}
}
