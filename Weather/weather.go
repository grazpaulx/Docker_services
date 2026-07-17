package main

import (
	"fmt"
	"net/http"
)

func weather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		city = "Unknown"
	}

	fmt.Fprintf(w, "Weather in %s: Sunny, 30°C", city)
}

func main() {
	http.HandleFunc("/weather", weather)

	fmt.Println("Weather Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
