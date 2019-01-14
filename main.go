package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/SimePel/asu-monitoring/proxy"
)

var (
	t = template.Must(template.New("T").ParseFiles("templates/index.html"))
)

func main() {
	server := &http.Server{Addr: ":8080", Handler: nil}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/proxy", proxyHandler)

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
