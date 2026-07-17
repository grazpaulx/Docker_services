package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func checkNumber(w http.ResponseWriter, r *http.Request) {
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

	if num%2 == 0 {
		fmt.Fprintf(w, "%d is Even", num)
	} else {
		fmt.Fprintf(w, "%d is Odd", num)
	}
}

func main() {
	http.HandleFunc("/check", checkNumber)

	fmt.Println("Even/Odd Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
