package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"path"
)


func main() {
	routes()
}

func routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":1337", r))
}

// serves index file
func HomeHandler(w http.ResponseWriter, r*http.Request) {
	// set header
	p := path.Dir("./public/views/index.html")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

