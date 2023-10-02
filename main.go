package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get the port from an environment variable or use a default value (e.g., 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the HTTP server using the specified port
	http.HandleFunc("/", EchoHandler)
	fmt.Printf("Server listening on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

// EchoHandler echoes back the request's URL path
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the entire URL path from the request, including query parameters
	fullPath := r.URL.String()

	// Write the full URL path back as the response
	fmt.Fprintf(w, "You requested: %s", fullPath)
}
