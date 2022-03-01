package main

import (
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/infrastructure/router"

	"net/http"
)

func main() {
	mysql.Connect()
	muxRouter := router.SetUp()
	err := http.ListenAndServe(":8888", muxRouter)
	if err != nil {
		panic(err)
	}
}
