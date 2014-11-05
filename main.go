package main

import (
	"github.com/dyerw/nonotevenmusic/api"
	"github.com/gorilla/mux"
	"github.com/jmcvetta/neoism"
	"net/http"
)

var DBConn *neoism.Database

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("home"))
}

func AlbumHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	json := ""
	if r.Method == "GET" {
		json = api.GetAlbum(r.URL.Query(), DBConn)
	}
	if r.Method == "PUT" {
		json = api.PutAlbum(r.FormValue("data"), DBConn)
	}

	rw.Write([]byte(json))
}

func main() {
	// Connect to database
	DBConn, _ = neoism.Connect("http://localhost:7474/db/data")

	r := mux.NewRouter()

	// Map url routes to functions
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/album", AlbumHandler)
	http.Handle("/", r)

	// Start server
	http.ListenAndServe(":3000", nil)
}
