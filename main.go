package main

import (
	"fmt"
	"net/http"
)

// there is the case where we use http.HandleFunc()
// it act like a wrapper
// it takes the path and the function and use the handle funct
// after adabting the function to the handler
// ////////////////////////////////////////////////////////////////////////////
type Wrapper struct {
	F func(http.ResponseWriter, *http.Request)
}

func (W *Wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	W.F(w, r)
}

func WrapperFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	instance := Wrapper{}
	instance.F = f
	http.Handle(path, &instance)
}

// /////////////////////////////////////////////////////////////////////////////
// this is the function that is going to be wrapped
func Example(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is function got adapted to the handler and rejesterd to the default multiplerxer")
}

// /////////////////////////////////////////////////////////////////////////////
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
	///////////////////////////////////
	// this function acts like the http.HandleFunc
	WrapperFunc("/path", Example)
	/////////////////////////////////
	server.ListenAndServe()
}
// visit localt host /path and you can see the prompt
// visit localhost:8080/hello or localhost:8080/world on your browser
