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
	godotenv.Load()
	recipientEmail := ""
	recipientName := "Contact"

	if r.Method == "POST" {
		msg := models.Reqbody{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Question: r.FormValue("questions"),
			Message:  r.FormValue("Message"),
		}

		switch msg.Question {
		case "Events":
			recipientEmail = "bporter@devonfarm.xyz"
			recipientName = "Events"
		case "Membership":
			recipientEmail = "lloeber@devonfarm.xyz"
			recipientName = "Membership"
		case "Language":
			recipientEmail = "mevans@snohomishtribe.org"
			recipientName = "Language"
		// case "Programs":
		// 	recipientEmail = "mevans@snohomishtribe.org"
		// 	recipientName = "Programs"
		default:
			recipientEmail = "contact@devonfarm.xyz"
		}

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
		// sends the email
		from := gomap.NewAddress(msg.Name, os.Getenv("EMAIL_ACCOUNT"))
		to := gomap.NewAddress(fmt.Sprintf("Snohomish Tribe %s", recipientName), recipientEmail) // Email subject title

		if err := mail.SendEmail(
			gomap.NewAddresses(from),
			gomap.NewAddresses(to),
			fmt.Sprintf("Contact Page Question: %s", msg.Question),
			fmt.Sprintf("From %s \n\n %s", msg.Email, msg.Message), // Email message
			false,
		); err != nil {
			log.Fatal(err, " line 76")
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
