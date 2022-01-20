package models

import (
	"net/http"
	"pack/db"
	"strconv"
)

type Clothing struct {
	Id           int64
	Name         string
	Breed        string
	Cost         float64
	CustomerCost float64
	Existence    int64
}

//Construct of Clothing
func buildClothing(name, breed string, cost, customercost float64, existence int64) Clothing {
	clothes := Clothing{
		Name:         name,
		Breed:        breed,
		Cost:         cost,
		CustomerCost: customercost,
		Existence:    existence,
	}
	return clothes
}

//Func to catch and to insert clothing in the Database
func InsertClothing(rw http.ResponseWriter, r *http.Request) {
	connection := db.OpenConnection()

	if r.Method == "POST" {

		name := r.FormValue("name")
		breed := r.FormValue("breed")

		cost, err := strconv.ParseFloat(r.FormValue("cost"), 64)
		if err != nil {
			panic(err.Error())
		}

		customerCost, err := strconv.ParseFloat(r.FormValue("customerCost"), 64)
		if err != nil {
			panic(err.Error())
		}
		existence, err := strconv.ParseInt(r.FormValue("existence"), 0, 64)
		if err != nil {
			panic(err.Error())
		}

		sql := "INSERT INTO productos (name, breed, cost, customerCost, existence) values (?,?,?,?,?)"
		responsePrepare, err := connection.Prepare(sql)
		if err != nil {
			panic(err.Error())
		}

		clothes := buildClothing(name, breed, cost, customerCost, existence)

		result, err := responsePrepare.Exec(clothes.Name, clothes.Breed, clothes.Cost, clothes.CustomerCost, clothes.Existence)

		if err != nil {
			panic(err.Error())
		}

		clothes.Id, err = result.LastInsertId()

		if err != nil {
			panic(err.Error())
		}

		defer responsePrepare.Close()
	}

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)

}
