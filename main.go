package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"path"
	"fmt"
)

// main func
func main() {
	routes()
}

// routes
func routes() {
	// init router
	r := mux.NewRouter()
	// index route
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/share", ShareHandler)
	r.HandleFunc("/download/", DownloadHandler)

	// start server on port 1337
	log.Fatal(http.ListenAndServe(":1337", r))
}

// serves index file
func HomeHandler(w http.ResponseWriter, r*http.Request) {
	p := path.Dir("./public/views/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

// get shared files
func ShareHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "ParseForm() err: %v", err)
			return
		}
		log.Println(r.FormValue("name"))
		http.Redirect(w, r, "/download/", http.StatusMovedPermanently)
	}
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("./public/views/share.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

