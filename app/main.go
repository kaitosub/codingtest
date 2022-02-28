package main

import (
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/infrastructure/router"

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
