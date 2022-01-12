package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sukanmi.com/lenslocked/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
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

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/faq", faqHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
