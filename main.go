package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var build_id, build_time = "-1.0", "today"

// Reference HTML templates
var siteTemplate = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	var port string
	// Make port configurable
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8000"
	}

	// Check tags
	fmt.Println("Build time: " + build_time)

	// Server logic (TODO - break this into a series of discrete functions, log requests)
	mux := http.NewServeMux()

	// Basic logging
	log.Printf("Starting server version: %v on port: %v", build_id, port)

	mux.HandleFunc("/", HandleGet)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	siteTemplate.Execute(w, nil)
}
