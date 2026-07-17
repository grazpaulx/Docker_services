package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func square(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("num")

	if numStr == "" {
		http.Error(w, "Please provide a number using ?num=", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := num * num

	fmt.Fprintf(w, "Square of %d is %d", num, result)
}

func main() {
	http.HandleFunc("/square", square)

	fmt.Println("Square Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
