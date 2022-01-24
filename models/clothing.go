package models

import (
	"fmt"
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

type Clothes []Clothing

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

	http.Redirect(rw, r, "/seeClothing", http.StatusMovedPermanently)

}

//FUNC TO LIST THE CLOTHIN
func ListClothing() Clothes {
	connection := db.OpenConnection()
	sql := "SELECT * FROM productos"
	rowsQueryClothing, err := connection.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	clothing := Clothes{}
	for rowsQueryClothing.Next() {
		clothes := Clothing{}
		rowsQueryClothing.Scan(&clothes.Id, &clothes.Name, &clothes.Breed, &clothes.Cost, &clothes.CustomerCost, &clothes.Existence)

		clothing = append(clothing, clothes)
	}

	return clothing
}

//FUNC TO SEARCH CLORHING BY  NAME
func SearchClothingbyName(r *http.Request) (clothing Clothes) {
	conexion := db.OpenConnection()

	if r.Method == "POST" {
		nameProduct := r.FormValue("search")
		sql := "SELECT * FROM productos where name=?"

		rowsQueryClothing, err := conexion.Query(sql, nameProduct)
		clothing = Clothes{}

		if err != nil {
			panic(err.Error())
		}
		for rowsQueryClothing.Next() {
			clothes := Clothing{}
			rowsQueryClothing.Scan(&clothes.Id, &clothes.Name, &clothes.Breed, &clothes.Cost, &clothes.CustomerCost, &clothes.Existence)

			clothing = append(clothing, clothes)
			fmt.Println(clothing)

		}
	}

	return
}

//FUCT TO DELETE A SPECIFIC CLOHING BY ID
func DeleteClothing(rw http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	connection := db.OpenConnection()
	sql := "DELETE FROM productos where idProducts=?"
	_, err := connection.Exec(sql, idProduct)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Producto eliminado")
	http.Redirect(rw, r, "/seeClothing", http.StatusMovedPermanently)
}

//FUNCS TO UPDATE A PRODUCT
func UpdatedClothing(r *http.Request) Clothing {
	idProduct := r.URL.Query().Get("id")
	connection := db.OpenConnection()
	sql := "SELECT * FROM productos where idProducts=?"
	productQueryClothing, err := connection.Query(sql, idProduct)

	if err != nil {
		panic(err.Error())
	}
	clothes := Clothing{}

	for productQueryClothing.Next() {
		productQueryClothing.Scan(&clothes.Id, &clothes.Name, &clothes.Breed, &clothes.Cost, &clothes.CustomerCost, &clothes.Existence)
	}
	return clothes
}

func UpdatedClothes(rw http.ResponseWriter, r *http.Request) {
	connection := db.OpenConnection()

	if r.Method == "POST" {

		idProduct := r.FormValue("id")
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

		sql := "UPDATE productos SET name=?, breed=?, cost=?, customerCost=?, existence=? where idProducts=?"
		responsePrepare, err := connection.Prepare(sql)
		if err != nil {
			panic(err.Error())
		}

		clothes := buildClothing(name, breed, cost, customerCost, existence)
		_, err = responsePrepare.Exec(&clothes.Name, &clothes.Breed, &clothes.Cost, &clothes.CustomerCost, &clothes.Existence, idProduct)
		fmt.Println("registro actualizado")
		if err != nil {
			panic(err.Error())
		}

		defer responsePrepare.Close()

	}
	http.Redirect(rw, r, "/seeClothing", http.StatusMovedPermanently)

}
