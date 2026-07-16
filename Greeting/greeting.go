package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s! Welcome to the Greeting Service.", name)
}

func main() {
	http.HandleFunc("/greet", greet)
	fmt.Println("Greeting Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
