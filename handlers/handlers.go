package handlers

import (
	"fmt"
	"html/template"

	"log"
	"net/http"
	"os"

	"github.com/cwinters8/gomap"
	"github.com/danielekpark/models"
	"github.com/joho/godotenv"
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
	recipient := ""
	if r.Method == "POST" {
		godotenv.Load()
		msg := models.Reqbody{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Question: r.FormValue("questions"),
			Message:  r.FormValue("message"),
		}

		switch msg.Question {
		case "events":
			recipient = " bporter@snohomishtribe.org"
		case "membership":
			recipient = "lloeber@snohomishtribe.org"
		default:
			recipient = "contact@snohomishtribe.org"
		}

		fmt.Println(recipient)
		//Start here
		mail, err := gomap.NewClient(
			"https://api.fastmail.com/jmap/session",
			os.Getenv("BEARER_TOKEN"),
			gomap.DefaultDrafts,
			gomap.DefaultSent,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mail)
		// send an email
		// from := gomap.NewAddress(msg.Name, msg.Email)
		from := gomap.NewAddress(msg.Name, "daniel@devonfarm.xyz")
		to := gomap.NewAddress("Snohomish Tribe Guest user", "daniel@devonfarm.xyz")

		if err := mail.SendEmail(
			gomap.NewAddresses(from),
			gomap.NewAddresses(to),
			"Contact Form Question",
			fmt.Sprintf("From %s %s", msg.Email, msg.Message),
			false,
		); err != nil {
			log.Fatal(err, " line 80")
		}

		w.WriteHeader(http.StatusOK)
		return
	}

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
