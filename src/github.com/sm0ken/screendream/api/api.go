package api

import (
	"fmt"
	"net/http"
)




func SpinUp() {

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "403")
	})

	http.HandleFunc("/api/togglePause", func(w http.ResponseWriter, r *http.Request) {
	})	

	http.HandleFunc("/api/next", func(w http.ResponseWriter, r *http.Request) {
	})	

	http.HandleFunc("/api/prev", func(w http.ResponseWriter, r *http.Request) {
	})	


	http.HandleFunc("/api/current", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/api/posturl", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/api/fetchplaylist", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/api/jumpsong", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/api/jumptime", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/api/clearplaylist", func(w http.ResponseWriter, r *http.Request) {
	})

	server := http.server{
		Addr: ":8080"
	}

	fs := http.FileServer(http.Dir("static/"))	// create file handler to serve static resources
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", fs)

	server.ListenAndServe()
}
