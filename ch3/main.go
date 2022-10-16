package main

import (
	"fmt"
	"net/http"
)

type HelloHundler struct{}

func (h *HelloHundler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHundler struct{}

func (h *WorldHundler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	helloHandler := HelloHundler{}
	worldHandler := WorldHundler{}
	// starting up the server
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &helloHandler)
	http.Handle("/world", &worldHandler)
	server.ListenAndServe()
}
