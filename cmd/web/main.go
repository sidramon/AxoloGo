package web

import (
	"fmt"
	"net/http"

	"github.com/sidramon/AxoloGo/internal/handlers"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println("(http://localhost:3000) - Server started on port", port)
	http.ListenAndServe(port, nil)
}