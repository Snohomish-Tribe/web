package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/snohomishtribe/pkg/jmap"
	"github.com/snohomishtribe/pkg/models"
)

const RECIPIENT_EMAIL_DOMAIN = "snohomishtribe.org"

func checkPath(expectedPath string, w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != expectedPath {
		tmpl, err := template.ParseFiles("static/templates/404.html", "static/templates/main.layout.html")
		if err != nil {
			return fmt.Errorf("failed to parse 404 template: %w", err)
		}
		if err := tmpl.Execute(w, nil); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}
	return nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/about", w, r); err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("static/templates/about.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/contact", w, r); err != nil {
		log.Printf("Error checking path: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	recipientEmail := ""
	recipientName := "Contact"

	if r.Method == "POST" {
		msg := models.Reqbody{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Question: r.FormValue("questions"),
			Message:  r.FormValue("Message"),
		}

		recipientPrefix := ""

		switch msg.Question {
		case "Events":
			recipientPrefix = "kvansenus"
			recipientName = "Events"
		case "Membership":
			recipientPrefix = "lloeber"
			recipientName = "Membership"
		case "Language":
			recipientPrefix = "mevans"
			recipientName = "Language"
		case "Programs":
			recipientPrefix = "lceniceros"
			recipientName = "Programs"
		default:
			recipientPrefix = "contact"
		}

		recipientEmail = fmt.Sprintf("%s@%s", recipientPrefix, RECIPIENT_EMAIL_DOMAIN)

		// Create JMAP client
		client, err := jmap.NewClient(
			"https://api.fastmail.com/jmap/session",
			os.Getenv("FASTMAIL_TOKEN"),
		)
		if err != nil {
			log.Printf("Failed to create JMAP client: %v\n", err)
			http.Error(w, "Failed to send email. Please try again later.", http.StatusInternalServerError)
			return
		}

		// Send the email
		from := jmap.NewAddress(msg.Name, os.Getenv("SENDER_EMAIL"))
		to := []jmap.Address{
			jmap.NewAddress(fmt.Sprintf("Snohomish Tribe %s", recipientName), recipientEmail),
		}

		if err := client.SendEmail(
			from,
			to,
			fmt.Sprintf("Contact Page Question: %s", msg.Question),
			fmt.Sprintf("From %s \n\n %s", msg.Email, msg.Message),
			false,
		); err != nil {
			log.Printf("Failed to send email: %v\n", err)
			http.Error(w, "Failed to send email. Please try again later.", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("static/templates/success.html", "static/templates/main.layout.html")
		if err != nil {
			log.Printf("Failed to parse success template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to execute success template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		return
	}

	tmpl, err := template.ParseFiles("static/templates/contact.html", "static/templates/main.layout.html")
	if err != nil {
		log.Printf("Failed to parse contact template: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Failed to execute contact template: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Credits(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/credits", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/credits.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/events", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/events.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Language(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/language-learning", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/language-learning.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Membership(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/membership", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/membership.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Programs(w http.ResponseWriter, r *http.Request) {
	if err := checkPath("/programs", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/programs.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}
