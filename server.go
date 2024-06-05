package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handles form error
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	// Return successful message that also show the user inputs
	fmt.Fprintf(w, "Request POST successfully:\n\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	// fmt.Fprintf(w, "name: %s, address: %s", name, address)
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s", address)
}

// Function for our request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Handles the request method if it's different from GET
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}

	// Return response for /hello route
	fmt.Fprintf(w, "Hello World!")
}

// The routes function
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
