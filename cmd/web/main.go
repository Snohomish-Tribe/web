package main

import (
	"net/http"

	"github.com/snohomishtribe/pkg/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/credits", handlers.Credits)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/language-learning", handlers.Language)
	http.HandleFunc("/membership", handlers.Membership)
	http.HandleFunc("/programs", handlers.Programs)

	http.ListenAndServe(":3000", nil)
}
