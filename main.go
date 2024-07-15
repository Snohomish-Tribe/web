package main

import (
	"net/http"

	"github.com/danielekpark/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)

	http.ListenAndServe(":3000", nil)
}
