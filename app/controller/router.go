package controller

import (
	"net/http"
)

type Router interface {
	HandleTransactionsRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TransactionController
}

func NewRouter(tc TransactionController) Router {
	return &router{tc}
}

func (ro *router) HandleTransactionsRequest(w http.ResponseWriter, r *http.Request) {
	//switch r.Method {
	//case "POST":
	ro.tc.PostTransaction(w, r)
	//default:
	//	w.WriteHeader(405)
	//}
}
