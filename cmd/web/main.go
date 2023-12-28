package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TelvoWarrior/bmwawg/pkg/config"
	"github.com/TelvoWarrior/bmwawg/pkg/handlers"
	"github.com/TelvoWarrior/bmwawg/pkg/render"
)

const portNumber = ":8080"

// main is the main function when all it begins
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port: %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
