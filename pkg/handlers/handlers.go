package handlers

import (
	"fmt"
	"html/template"

	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cwinters8/gomap"
	"github.com/joho/godotenv"
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

func FetchEvents(url string, token string) models.EventsPageData {
	godotenv.Load()
	var strapiRes models.EventsPageData

	client := &http.Client{Timeout: time.Second * 10}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("There was no response from Strapi")
			}
		}()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read response")
	} else {
		json.Unmarshal(body, &strapiRes)
		fmt.Println(strapiRes)
	}

	defer res.Body.Close()

	return strapiRes
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
		log.Fatal(err)
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
			recipientPrefix = "bporter"
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

		mail, err := gomap.NewClient(
			"https://api.fastmail.com/jmap/session",
			os.Getenv("FASTMAIL_TOKEN"),
			gomap.DefaultDrafts,
			gomap.DefaultSent,
		)
		if err != nil {
			log.Fatal(err)
		}
		// sends the email
		from := gomap.NewAddress(msg.Name, os.Getenv("SENDER_EMAIL"))
		to := gomap.NewAddress(fmt.Sprintf("Snohomish Tribe %s", recipientName), recipientEmail)

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
	if err := checkPath("/credits", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/credits.html", "static/templates/main.layout.html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Failed to parse template ", err)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("PROD_EVENTS_BEARER_TOKEN")
	url := os.Getenv("PROD_EVENTS_ENDPOINT")
	ParsedData := FetchEvents(url, token)

	if err := checkPath("/events", w, r); err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("static/templates/events.html", "static/templates/main.layout.html")

	// if err := tmpl.Execute(w, nil); err != nil {
	// 	log.Fatal("Failed to parse template ", err)
	// }

	if err := tmpl.Execute(w, ParsedData); err != nil {
		log.Fatal("Failed to parse template", err)
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
