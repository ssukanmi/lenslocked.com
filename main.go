package main

import (
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ip := r.RemoteAddr
	dateTime := time.Now()
	fmt.Fprintf(w, "<h1>Welcome to GoLang</h1>Your IP: %s<br>Date: %s", ip, dateTime.Format(time.ANSIC))
}

func main() {
	fmt.Println("Server started!!!")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
