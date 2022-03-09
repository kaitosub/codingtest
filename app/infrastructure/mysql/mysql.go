package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	DBMS := "mysql"
	//dbInfo := fmt.Sprintf(
	//	"%s:%s@%s/%s?parseTime=true",
	//	"go_test",
	//	"root",
	//	"tcp@127.0.0.1:3306",
	//	"go_database",
	//)
	db, err := sql.Open(DBMS, "root@tcp(127.0.0.1)/codetest")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	DB = db
}
