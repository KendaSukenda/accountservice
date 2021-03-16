package repository

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

func GetDBConnection() *sql.DB {
	dbName := "go_ms_test"
	fmt.Println("Conn info:", dbName)
	db, err := sql.Open("mysql",
		"root:sukenda@tcp(127.0.0.1:3306)/go_ms_test")
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	return db
}
