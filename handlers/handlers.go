package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	// "strings"
	// "github.com/danielekpark/models"
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

	if r.Method == "POST" {
		// msg := models.Reqbody{
		// 	Name:     r.FormValue("name"),
		// 	Email:    r.FormValue("email"),
		// 	Question: r.FormValue("questions"),
		// 	Message:  r.FormValue("message"),
		// }

		fmt.Println(r.FormValue("name"))
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
	// events go to bporter@snohomishtribe.org
	//- membership goes to lloeber@snohomishtribe.org
	//- general goes to contact@snohomishtribe.org

}

//https://pkg.go.dev/github.com/cwinters8/gomap#example-Client.SendEmail
// https://www.kirandev.com/http-post-golang
