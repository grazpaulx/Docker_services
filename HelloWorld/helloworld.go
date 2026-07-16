package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Welcome to Docker with Go.")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Hello World Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
