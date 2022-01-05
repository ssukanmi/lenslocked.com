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
	fmt.Fprint(w, "<h1>Welcome to Homepage</h1>")
	fmt.Fprintf(w, "Your IP: %s<br/>", ip)
	fmt.Fprintf(w, "Date: %s<br/>", dateTime.Format((time.ANSIC)))
	fmt.Fprint(w, "<a href=\"/contact\">Contact page</a><br/>")
	fmt.Fprint(w, "<a href=\"/faq\">FAQ page</a><br/>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.<br/>")
	fmt.Fprint(w, "<a href=\"/\">Go back home</a><br/>")
	fmt.Fprint(w, "<a href=\"/faq\">FAQ page</a><br/>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is the content of the FAQ page.</h1>")
	fmt.Fprint(w, "<a href=\"/\">Go back home</a><br/>")
	fmt.Fprint(w, "<a href=\"/contact\">Contact page</a><br/>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")
	fmt.Fprint(w, "<a href=\"/\">Go back home</a><br/>")
}

func main() {
	fmt.Println("Server started!!!")
	fmt.Println("Testing cli branching")

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/faq", faqHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
