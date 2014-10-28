package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("home"))
}

func AlbumHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	json := ""
	if r.Method == "POST" {
		json = GetAlbum(r.FormValue("data"))
	}
	if r.Method == "PUT" {
		json = PutAlbum(r.FormValue("data"))
	}

	rw.Write([]byte(json))
}

func main() {
	r := mux.NewRouter()

	// Map url routes to functions
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/album", AlbumHandler)
	http.Handle("/", r)

	// Start server
	http.ListenAndServe(":3000", nil)
}
