package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home page.")
}

func main() {
	
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8080", nil)

}