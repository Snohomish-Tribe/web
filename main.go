package main

import (
	"net/http"

	"github.com/danielekpark/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//Get Requests
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/credits", handlers.Credits)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/events", handlers.Events)

	// Post request

	http.ListenAndServe(":3000", nil)
}
