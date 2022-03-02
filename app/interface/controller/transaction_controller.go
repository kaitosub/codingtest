package controller

import (
	"encoding/json"
	"github.com/kaitosub/codingtest/app/entity/model"
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/interface/database"
	"github.com/kaitosub/codingtest/app/usecase"
	"log"
	"net/http"
	"strconv"
)

type TransactionController struct {
	interactor usecase.TransactionInteractorInterface
}

func NewTransactionController() TransactionControllerInterface {
	return &TransactionController{interactor: usecase.NewTransactionInteractor(&database.TransactionRepository{})}
}

type TransactionControllerInterface interface {
	GetTransactions(w http.ResponseWriter, r *http.Request)
}

var amountLimit = 1000

func (tr *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		log.Fatalln("33:", err)
		return
	}
	var transactionRequest model.TransactionRequest
	err = json.Unmarshal(body, &transactionRequest)
	if err != nil {
		log.Fatalln("39:", err)
		return
	}

	transaction := model.TransactionEntity{UserID: transactionRequest.UserId, Amount: transactionRequest.Amount, Description: transactionRequest.Description}

	//var amount int
	//if err := mysql.DB.QueryRow(
	//	"select sum(amount) from transactions where user_id=?",
	//	ctxUser.ID,
	//).Scan(&amount); err != nil {
	//	log.Fatal(err)
	//}
	//if amount > amountLimit {
	//	log.Printf("amount %d over the amountLimit %d", amount, amountLimit)
	//}
	//
	//if amount+transaction.Amount > amountLimit {
	//	log.Print("error")
	//	return
	//}

	id, err := InsertTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("61:", err)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("77:", err, r)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func InsertTransaction(transaction model.TransactionEntity) (id int, err error) {
	_, err = mysql.DB.Exec("INSERT INTO transaction (user_id, amount, description) VALUES (?, ?, ?)", transaction.UserID, transaction.Amount, transaction.Description)
	return id, err
}
