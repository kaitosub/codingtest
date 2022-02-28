package router

import (
	"github.com/gorilla/mux"
	"github.com/mfkessai/codetest-docker/app/interface/controller"
)

func SetUp() *mux.Router {
	r := mux.NewRouter()

	transactionController := controller.NewTransactionController()
	r.HandleFunc("/transactions", transactionController.GetTransactions).Methods("POST")

	return r
}
