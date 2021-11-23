package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my awesome site")
}

func main() {
	fmt.Println("Server started!!!")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
