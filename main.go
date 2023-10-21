package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "this is GET Method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Creates a file server handler that will serve files from the "./static" directory.
	fileServer := http.FileServer(http.Dir("./static"))

	// registers the file server handler to handle requests for the root URL ("/").
	http.Handle("/", fileServer)

	//registers a handler function formHandler to handle requests for the "/form" URL and hello
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Sever at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // if error check pattern
		log.Fatal(err)
	}
}
