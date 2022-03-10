package main

import (
	"kaitosub/app/controller"
	"kaitosub/app/model/repository"
	"net/http"
)

var tr = repository.NewTransactionRepository()
var tc = controller.NewTransactionController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8888",
	}
	http.HandleFunc("/transactions/", ro.HandleTransactionsRequest)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
