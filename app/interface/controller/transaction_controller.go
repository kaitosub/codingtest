package controller

import (
	"encoding/json"
	"github.com/mfkessai/codetest-docker/app/entity/model"
	"github.com/mfkessai/codetest-docker/app/infrastructure/mysql"
	"github.com/mfkessai/codetest-docker/app/interface/database"
	"github.com/mfkessai/codetest-docker/app/usecase"
	"github.com/mfkessai/codetest-docker/app/util/ctx"
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
	w.Header().Set("Content-Type", "application/json")
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var transactionRequest model.Transaction
	json.Unmarshal(body, &transactionRequest)

	ctxUser, err := ctx.GetCtxUser(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	transanction := model.Transaction{UserId: ctxUser.ID, Amount: transactionRequest.Amount, Description: transactionRequest.Description}

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

	if amount+transanction.Amount > amountLimit {
		log.Print("error")
		return
	} else {
		id, err := InsertTransaction(transanction)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//err = json.NewEncoder(w).Encode(transactions)
		if err != nil {
			//logger.ErrDump(err, r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
		w.WriteHeader(201)
	}
}

func InsertTransaction(transaction model.Transaction) (id int, err error) {
	_, err = mysql.DB.Exec("INSERT INTO transaction (user_id, amount, description) VALUES (?, ?, ?)", transaction.UserId, transaction.Amount, transaction.Description)
	return id, err
}
