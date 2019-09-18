package main

import (
	"github.com/sm0ken/screendream/video"
	"github.com/sm0ken/screendream/downloader"
	"fmt"
	"log"
	"net/http"
	"html/template"
)

var Playlist *video.Playlist = video.NewPlaylist()


type rootHandler struct {
	clients int
}



func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Root served ", h.clients, " times")

    tmpl := template.Must(template.ParseFiles("./src/github.com/sm0ken/screendream/api/index.html"))
    tmpl.Execute(w, Playlist)

    http.ListenAndServe(":80", nil)

	//http.ServeFile(w, r, "./static/index.html")
}


type playHandler struct {
}

func (h *playHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type nextHandler struct {
}

func (h *nextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type prevHandler struct {
}

func (h *prevHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type currentHandler struct {
}

func (h *currentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}


type postHandler struct {
}

func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {


    if err := r.ParseForm(); err != nil {
        fmt.Println(w, "ParseForm() err: %v", err)
        return
    }
    url := r.FormValue("url")

    vid := downloader.Download_video(url)
	Playlist.AddVideo(vid)

	tmpl := template.Must(template.ParseFiles("./src/github.com/sm0ken/screendream/api/index.html"))
    tmpl.Execute(w, Playlist)

    http.ListenAndServe(":80", nil)
}


type fetchHandler struct {
}

func (h *fetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type jumpHandler struct {
}

func (h *jumpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type seekHandler struct {
}

func (h *seekHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type clearHandler struct {
}

func (h *clearHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type forbiddenHandler struct {
}

func (h *forbiddenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403 - Forbidden", 403)
}

func SpinUp() {
	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/", new(rootHandler))

	http.Handle("/api/", new(forbiddenHandler))
	http.Handle("/src/", new(forbiddenHandler))

	http.Handle("/api/play", new(playHandler))
	http.Handle("/api/next", new(nextHandler))
	http.Handle("/api/prev", new(prevHandler))
	http.Handle("/api/current", new(currentHandler))
	http.Handle("/api/post", new(postHandler))
	http.Handle("/api/fetch", new(fetchHandler))
	http.Handle("/api/jump", new(jumpHandler))
	http.Handle("/api/seek", new(seekHandler))
	http.Handle("/api/clear", new(clearHandler))

	fs := http.FileServer(http.Dir("static/")) // fs is a httphandler that serves static files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Print(server.ListenAndServe())
}

func main() {
	SpinUp()
}
