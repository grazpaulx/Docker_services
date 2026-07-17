package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func factorial(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("num")

	if numStr == "" {
		http.Error(w, "Please provide a number using ?num=", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 {
		http.Error(w, "Please enter a valid positive integer", http.StatusBadRequest)
		return
	}

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	fmt.Fprintf(w, "Factorial of %d is %d", num, result)
}

func main() {
	http.HandleFunc("/factorial", factorial)

	fmt.Println("Factorial Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
