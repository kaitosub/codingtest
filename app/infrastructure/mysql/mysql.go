package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	DBMS := "mysql"
	db, err := sql.Open(DBMS, "root@tcp(127.0.0.1)/codetest")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	DB = db
}
