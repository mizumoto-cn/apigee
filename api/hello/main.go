package main

// This file defines a simple "Hello, world!" HTTP Cloud Function.
// It reads a "name" parameter from the request and writes a response.

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "<p>Hello, %s!</p>", name)
}

func main() {
	port := "10086"
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
