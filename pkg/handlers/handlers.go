package handlers

import (
	"fmt"
	"html/template"

	"log"
	"net/http"
	"os"

	"github.com/cwinters8/gomap"
	"github.com/joho/godotenv"
	"github.com/snohomishtribe/pkg/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/index.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/about.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	// recipient := ""
	godotenv.Load()

	if r.Method == "POST" {
		msg := models.Reqbody{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Question: r.FormValue("questions"),
			Message:  r.FormValue("message"),
		}

		// switch msg.Question {
		// case "events":
		// 	recipient = " bporter@snohomishtribe.org"
		// case "membership":
		// 	recipient = "lloeber@snohomishtribe.org"
		// default:
		// 	recipient = "contact@snohomishtribe.org"
		// }
		// fmt.Println(recipient)

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
		from := gomap.NewAddress(msg.Name, os.Getenv("EMAIL_ACCOUNT"))
		to := gomap.NewAddress("Snohomish Tribe Website Contact Form Question", os.Getenv("EMAIL_ACCOUNT")) // Email subject title

		if err := mail.SendEmail(
			gomap.NewAddresses(from),
			gomap.NewAddresses(to),
			fmt.Sprintf("Contact Page Question: %s", msg.Question),
			fmt.Sprintf("From %s \n\n %s", msg.Email, msg.Message), // Message
			false,
		); err != nil {
			log.Fatal(err, " line 75")
		}

		tmpl, _ := template.ParseFiles("static/templates/success.html", "static/templates/main.layout.html")
		if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal("Failed to parse template ", err)
		}

		return
	}

	tmpl, _ := template.ParseFiles("static/templates/contact.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Credits(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/credits.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/events.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Language(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/language-learning.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Membership(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/membership.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Programs(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/programs.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}
