package main

import (
	"fmt"
	"net/http"

	"github.com/snohomishtribe/pkg/handlers"
)

const PORT = 3000

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
	http.HandleFunc("/not-found", handlers.NotFound)

	fmt.Printf("Listening on port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
