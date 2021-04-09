package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
	})

	http.ListenAndServe(":"+port, nil)
}
