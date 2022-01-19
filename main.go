package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("template/**/*.html"))

func main() {
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/seeClothing", ListClothing)
	//server
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}

func renderTemplate(rw http.ResponseWriter, name string) {
	temp.ExecuteTemplate(rw, name, nil)
}

func Index(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "index.html")
}

func ListClothing(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "seeClothing.html")
}
