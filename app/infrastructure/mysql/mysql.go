package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	db, err := sqlx.Connect("mysql", "root@tcp(127.0.0.1)/codetest")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	DB = db
}
