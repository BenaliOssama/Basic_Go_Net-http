package main

import "net/http"

func main() {
	// starting a web server
	// if i leave the port as default
	// then port 80 will be used
	// in that case you need to run the 
	// with sudo previlage
	http.ListenAndServe("", nil)
}
