package main

import (
	"fmt"
	"net/http"
	"time"
)

func currentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	fmt.Fprintf(w, "Current Date and Time: %s", now.Format("02-Jan-2006 03:04:05 PM"))
}

func main() {
	http.HandleFunc("/time", currentTime)

	fmt.Println("Current Time Service running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
