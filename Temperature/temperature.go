package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func convertTemperature(w http.ResponseWriter, r *http.Request) {
	celsiusStr := r.URL.Query().Get("c")

	if celsiusStr == "" {
		http.Error(w, "Please provide temperature using ?c=", http.StatusBadRequest)
		return
	}

	celsius, err := strconv.ParseFloat(celsiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Fprintf(w, "%.2f°C = %.2f°F", celsius, fahrenheit)
}

func main() {
	http.HandleFunc("/convert", convertTemperature)

	fmt.Println("Temperature Converter Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
