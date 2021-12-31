package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ip := r.RemoteAddr
	dateTime := time.Now()
	fmt.Fprintf(w, "<h1>Welcome to Homepage</h1>Your IP: %s<br>Date: %s<br>", ip, dateTime.Format(time.ANSIC))
}

func main() {
	fmt.Println("Server started!!!")

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
