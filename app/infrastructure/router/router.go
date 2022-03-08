package router

import (
	"github.com/gorilla/mux"
	"github.com/kaitosub/codingtest/app/interface/controller"
)

func SetUp() *mux.Router {
	r := mux.NewRouter()

	transactionController := controller.NewTransactionController()
	r.HandleFunc("/transactions/", transactionController.PostTransaction).Methods("POST")

	return r
}
