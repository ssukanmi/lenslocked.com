package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ip := r.RemoteAddr
	dateTime := time.Now()
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to Homepage</h1>Your IP: %s<br>Date: %s<br>", ip, dateTime.Format(time.ANSIC))
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>We could not find the page you were looking for :(</h1>")
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	fmt.Println("Server started!!!")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
