package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"path"
	"fmt"
	"bytes"
	"strings"
	"io"
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
	r.HandleFunc("/", home)
	r.HandleFunc("/upload", upload)
	r.HandleFunc("/download/", download)

	// start server on port 1337
	log.Fatal(http.ListenAndServe(":1337", r))
}

// serves index file
func home(w http.ResponseWriter, r*http.Request) {
	p := path.Dir("./static/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

// get shared files
func upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "ParseForm() err: %v", err)
			return
		}

		var Buf bytes.Buffer
		// in your case file would be fileupload
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		// Copy the file data to buffer
		io.Copy(&Buf, file)
		contents := Buf.String()
		fmt.Println(contents)
		Buf.Reset()
		return

		log.Println(r.FormValue("uploadfile"))
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, "./static/share.html")
}

