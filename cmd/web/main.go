package main

import (
	"fmt"
	"net/http"

	"github.com/TelvoWarrior/bmwawg/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function when all it begins
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port: %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
