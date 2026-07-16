package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculator(w http.ResponseWriter, r *http.Request) {

	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	op := r.URL.Query().Get("op")

	switch op {
	case "add":
		fmt.Fprintf(w, "Addition = %d", a+b)
	case "sub":
		fmt.Fprintf(w, "Subtraction = %d", a-b)
	case "mul":
		fmt.Fprintf(w, "Multiplication = %d", a*b)
	case "div":
		if b == 0 {
			fmt.Fprintf(w, "Cannot divide by zero")
			return
		}
		fmt.Fprintf(w, "Division = %d", a/b)
	default:
		fmt.Fprintf(w, "Invalid Operation")
	}
}

func main() {
	http.HandleFunc("/calculator", calculator)

	fmt.Println("Calculator Service Started...")
	fmt.Println("Open: http://localhost:8080/calculator?a=10&b=5&op=add")

	http.ListenAndServe(":8080", nil)
}
