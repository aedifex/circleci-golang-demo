package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
)

var build_id, build_time = "1234567", "117 AD"

// Reference HTML templates
var siteTemplate = template.Must(template.ParseFiles("templates/index.html"))

type App struct {
	// points to a mux router?
	Router *mux.Router
}

type AppVersion struct {
	Build_id   string
	Build_time string
}

func (a *App) Run() {
	a.Router = mux.NewRouter()
	// Map our functions to a handler
	a.Router.HandleFunc("/", HandleGet)
	a.Router.HandleFunc("/version", HandleVersion)
	a.Router.HandleFunc("/runtime", HandleRuntimeInfo)
}

func jsonIfy(element interface{}) ([]byte, error) {
	json, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	return json, nil
}

// Returns binary version in the form of SHA1 && compile time.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	app_version := AppVersion{Build_id: build_id, Build_time: build_time}
	payload, _ := jsonIfy(app_version)
	fmt.Fprintf(w, string(payload))
}

func main() {
	var port string
	// Make port configurable
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}

	// Check tags
	fmt.Println("Build time: " + build_time)

	a := App{}
	a.Run()
	log.Printf("Starting server version: %v on port: %v", build_id[0:7], port)
	err := http.ListenAndServe(port, a.Router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	siteTemplate.Execute(w, nil)
}

// Returns OS & ARCH info about the host.
func HandleRuntimeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
