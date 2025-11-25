package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/snohomishtribe/pkg/handlers"

	"github.com/joho/godotenv"
)

const PORT = 3000

func main() {
	godotenv.Load()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/credits", handlers.Credits)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/language-learning", handlers.Language)
	http.HandleFunc("/membership", handlers.Membership)
	http.HandleFunc("/programs", handlers.Programs)

	fmt.Printf("Listening on port %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}
}
