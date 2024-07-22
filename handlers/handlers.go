package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	// "strings"
	// "bytes"
	// "github.com/cwinters8/gomap"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

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

func Language(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/language-learning.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable to parse html")
	}
}

func Membership(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/membership.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable to parse html")
	}
}

func Programs(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/programs.html")

	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Unable to parse html")
	}
}

func SendEmail() {

}

// events go to bporter@snohomishtribe.org
//- membership goes to lloeber@snohomishtribe.org
//- general goes to contact@snohomishtribe.org
//https://pkg.go.dev/github.com/cwinters8/gomap#example-Client.SendEmail
// https://www.kirandev.com/http-post-golang
