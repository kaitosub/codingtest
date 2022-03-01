package controller

import (
	"encoding/json"
	"github.com/kaitosub/codingtest/app/entity/model"
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/interface/database"
	"github.com/kaitosub/codingtest/app/usecase"
	"github.com/kaitosub/codingtest/app/util/ctx"
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
	r.Body.Read(body)
	var transactionRequest model.Transaction
	json.Unmarshal(body, &transactionRequest)

	ctxUser, err := ctx.GetCtxUser(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	transaction := model.Transaction{UserId: ctxUser.ID, Amount: transactionRequest.Amount, Description: transactionRequest.Description}

	var amount int
	if err := mysql.DB.QueryRow(
		"select sum(amount) from transactions where user_id=?",
		ctxUser.ID,
	).Scan(&amount); err != nil {
		log.Fatal(err)
	}
	if amount > amountLimit {
		log.Printf("amount %d over the amountLimit %d", amount, amountLimit)
	}

	if amount+transaction.Amount > amountLimit {
		log.Print("error")
		return
	}

	id, err := InsertTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		//logger.ErrDump(err, r)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func InsertTransaction(transaction model.Transaction) (id int, err error) {
	_, err = mysql.DB.Exec("INSERT INTO transaction (user_id, amount, description) VALUES (?, ?, ?)", transaction.UserId, transaction.Amount, transaction.Description)
	return id, err
}
