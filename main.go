package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello from our simple web server")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Fprintf(w, "Form details:\n")
	fmt.Fprintf(w, "First name: %s\n", fname)
	fmt.Fprintf(w, "Last name: %s", lname)
}
