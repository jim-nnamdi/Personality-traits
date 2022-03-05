package main

import "database/sql"

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func DatabaseConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/personalitytest")
	ErrorCheck(err)
	return db
}

func main() {

}
