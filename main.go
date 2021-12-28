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
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to Homepage</h1>Your IP: %s<br>Date: %s<br>", ip, dateTime.Format(time.ANSIC))
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	}
}

func main() {
	fmt.Println("Server started!!!")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
