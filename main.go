package main

import "net/http"

func main() {
	// we creat an http server struct
	server := http.Server{
		// 127.0.0.1 = localhost
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	// this code now will run and hundle any request 
	// with 404 not found
	server.ListenAndServe()
}
