package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register request handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to kaylaconsta.com. You're currently visiting: %s\n", r.URL.Path)
	})

	// use absolute file paths, not relative nor Join
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listen for HTTP connections
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
