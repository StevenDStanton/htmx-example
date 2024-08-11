package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StevenDStanton/htmx-example/internal/handlers"
)

const port = ":8000"

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.GetBase)
	mux.HandleFunc("POST /contacts", handlers.CreateContact)

	fmt.Printf("Listening on Port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Could not listen on port %s: %v\n", port, err)
	}
}
