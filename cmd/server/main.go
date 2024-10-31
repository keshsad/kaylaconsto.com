package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to kaylaconsta.com. You're currently visiting: %s\n", r.URL.Path)
	})

	// listen for HTTP connections
	http.ListenAndServe(":42069", nil)
}
