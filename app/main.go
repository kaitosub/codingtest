package main

import (
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/infrastructure/router"
	"net/http"
)

//var tr = usecase.NewTransactionInteractor()
//var tc = controller.NewTransactionController(tr)
//var ro = router.NewRouter(tc)

func main() {
	//server := http.Server{
	//	Addr: ":8888",
	//}
	//http.HandleFunc("/transactions/", ro.HandleTransactionsRequest)
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	mysql.Connect()
	muxRouter := router.SetUp()
	err := http.ListenAndServe(":8888", muxRouter)
	if err != nil {
		//panic(err)
	}
}
