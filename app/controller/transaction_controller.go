package controller

import (
	"encoding/json"
	"github.com/kaitosub/codingtest/app/controller/dto"
	"github.com/kaitosub/codingtest/app/model/entity"
	"github.com/kaitosub/codingtest/app/model/repository"
	"net/http"
	"strconv"
)

type TransactionController interface {
	PostTransaction(w http.ResponseWriter, r *http.Request)
}

type transactionController struct {
	tr repository.TransactionRepository
}

func NewTransactionController(tr repository.TransactionRepository) TransactionController {
	return &transactionController{tr}
}

// TODOの追加
func (tc *transactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	// リクエストbodyのJSONをDTOにマッピング
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var transactionRequest dto.TransactionRequest
	json.Unmarshal(body, &transactionRequest)

	// DTOをTODOのEntityに変換
	transaction := entity.TransactionEntity{UserID: transactionRequest.UserID, Amount: transactionRequest.Amount, Description: transactionRequest.Description}

	// リポジトリの追加処理呼び出し
	id, err := tc.tr.InsertTransaction(transaction)
	if err != nil {
		w.WriteHeader(600)
		return
	}

	// LocationにリソースのPATHを設定し、ステータスコード２０１を返却
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}
