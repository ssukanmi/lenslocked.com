package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
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

	var err error
	homeTemplate, err = template.ParseFiles("views/home.go.html")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.go.html")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/faq", faqHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
