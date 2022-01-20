package main

import (
	"fmt"
	"log"
	"net/http"
	"pack/models"
	"text/template"
)

var temp = template.Must(template.ParseGlob("template/**/*.html"))

func main() {
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/seeClothing", ListClothing)
	mux.HandleFunc("/addClothing", addClothing)
	mux.HandleFunc("/insertClothing", models.InsertClothing)

	//server
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}

//Execute the template and send the struct to the html
func renderTemplate(rw http.ResponseWriter, name string) {
	temp.ExecuteTemplate(rw, name, nil)
}

func Index(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "index.html")
}

func ListClothing(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "seeClothing.html")
}

func addClothing(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "addClothing.html")
}
