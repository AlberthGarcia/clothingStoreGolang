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
	mux.HandleFunc("/addClothing", addClothing)
	mux.HandleFunc("/seeClothing", listClothing)
	mux.HandleFunc("/insertClothing", models.InsertClothing)
	mux.HandleFunc("/searchClothing", searchClothing)
	mux.HandleFunc("/updatedClothes", updatedClothing)
	mux.HandleFunc("/deleteClothing", models.DeleteClothing)
	mux.HandleFunc("/updatedClothing", models.UpdatedClothes)

	//server
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}

//Execute the template and send the struct to the html
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := temp.ExecuteTemplate(rw, name, data)
	if err != nil {
		panic(err.Error())
	}
}

func Index(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "index.html", nil)
}

func listClothing(rw http.ResponseWriter, r *http.Request) {
	list := models.ListClothing()
	renderTemplate(rw, "seeClothing.html", list)
}

func addClothing(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "addClothing.html", nil)
}

func searchClothing(rw http.ResponseWriter, r *http.Request) {
	listClothing := models.SearchClothingbyName(r)
	renderTemplate(rw, "searchClothing.html", listClothing)
}

func updatedClothing(rw http.ResponseWriter, r *http.Request) {
	clothes := models.UpdatedClothing(r)
	renderTemplate(rw, "updatedClothes.html", clothes)
}
