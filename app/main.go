package main

import (
	//"./infrastructure/mysql"
	"github.com/kaitosub/app/infrastructure/mysql"
	"github.com/kaitosub/app/infrastructure/router"
	//"./infrastructure/router"

	"net/http"
)

func main() {
	mysql.Connect()
	muxRouter := router.SetUp()
	err := http.ListenAndServe(":8080", muxRouter)
	if err != nil {
		panic(err)
	}
}
