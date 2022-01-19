package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const urlConnection = "root:@tcp(localhost:3306)/clothingStore"

var referenceDataBase *sql.DB

func OpenConnection() (connection *sql.DB) {

	connection, err := sql.Open("mysql", urlConnection)
	if err != nil {
		panic(err.Error())
	}

	referenceDataBase = connection
	return

}

func CloseConnection() {
	err := referenceDataBase.Close()
	if err != nil {
		panic(err.Error())
	}
}
