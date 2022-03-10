package repository

import (
	"database/sql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root@tcp(127.0.0.1)/codetest")

	if err != nil {
		panic(err)
	}
}
