package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	// Register request handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to kaylaconsta.com. You're currently visiting: %s\n", r.URL.Path)
	})

	// Set up the static directory
	staticDir := filepath.Join("..", "..", "web", "static")
	fmt.Println("Static directory:", staticDir) // Log the static directory path

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listen for HTTP connections
	if err := http.ListenAndServe(":42069", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
