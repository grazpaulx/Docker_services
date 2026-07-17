package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func randomNumber(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(100) + 1

	fmt.Fprintf(w, "Random Number: %d", number)
}

func main() {
	http.HandleFunc("/random", randomNumber)

	fmt.Println("Random Number Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
