package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/SimePel/asu-monitoring/proxy"
	"github.com/SimePel/asu-monitoring/web"
)

var (
	t = template.Must(template.New("T").ParseFiles("templates/index.html"))
)

func main() {
	server := &http.Server{Addr: ":8080", Handler: nil}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/proxy", proxyHandler)
	http.HandleFunc("/web", webHandler)

	server.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index", nil)
}

// Proxy struct is wrapper for json
type Proxy struct {
	Status bool `json:"status"`
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		Proxy{
			Status: proxy.Check(),
		})
}

// Web struct is wrapper for json
type Web struct {
	Status bool `json:"status"`
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		Web{
			Status: web.Check(),
		})
}
