package main

import (
	"fmt"
	"net/http"
)

// the handler we used previously handles all request with the same response
// if we want to handle diffrent request with different responses
// we need to make more handlers
type Hello struct{}

func (E *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type World struct{}

func (E *World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	// we define a handler multip handler
	hello := Hello{}
	world := World{}
	// in this case we use the default handler
	// to take care of the methods used to handle
	// different request
	server := http.Server{
		// 127.0.0.1 = localhost
		Addr: "127.0.0.1:8080",
	}
	// this code now will run and hundle any request
	// with 404 not found
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
// now visit localhost:8080/hello or localhost:8080/world on your browser 