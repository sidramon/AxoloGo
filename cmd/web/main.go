package main

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handlers.)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println("(http://localhost:3000) - Server started on port", port)
	http.ListenAndServe(port, nil)
}