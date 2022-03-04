package router

import (
	"github.com/kaitosub/codingtest/app/interface/controller"
	"net/http"
)

type Router interface {
	HandleTransactionsRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc controller.TransactionControllerInterface
}

func NewRouter(tc controller.TransactionControllerInterface) Router {
	return &router{tc}
}

func (ro *router) HandleTransactionsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ro.tc.PostTransaction(w, r)
	default:
		w.WriteHeader(405)
	}
}

//func SetUp() *mux.Router {

//r := mux.NewRouter()
//
//transactionController := controller.NewTransactionController()
//r.HandleFunc("/transactions/", transactionController.GetTransactions).Methods("POST")
//if err != nil {
//	return err
//}

//return r
//}

//hogehog
