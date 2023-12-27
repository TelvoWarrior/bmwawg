package main

import (
	"net/http"
)

const portNumber = ":8080"

// main is the main function when all it begins
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
