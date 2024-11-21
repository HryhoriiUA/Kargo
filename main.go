package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	http.HandleFunc("/", helloWorld)
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
