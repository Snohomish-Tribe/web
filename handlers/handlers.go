package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
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

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/contact.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable to parse html")
	}
}

func Credits(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/credits.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable parse html")
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/events.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable to parse html")
	}
}
