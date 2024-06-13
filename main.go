package main

import (
	"fmt"
	"net/http"
)

// what if we want to make my own hundler
// any type that has the method ServeHTTP(http.ResponseWriter, *http.Request)
// can be a handler
// EX:

type Example struct{}

func (E *Example) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is an example")
}

func main() {
	// we define a handler
	// this hundler will respond to any request by the same thing
	// "this is an example"
	handler := Example{}
	// we creat an http server struct
	server := http.Server{
		// 127.0.0.1 = localhost
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	// this code now will run and hundle any request
	// with 404 not found
	server.ListenAndServe()
}
