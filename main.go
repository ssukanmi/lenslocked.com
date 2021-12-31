package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ip := r.RemoteAddr
	dateTime := time.Now()
	fmt.Fprintf(w, "<h1>Welcome to Homepage</h1>Your IP: %s<br>Date: %s<br>", ip, dateTime.Format(time.ANSIC))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	fmt.Println("Server started!!!")

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
