package usecase

import (
	"github.com/kaitosub/codingtest/app/entity/model"
)

type TransactionInteractorInterface interface {
	PostTransaction(transaction model.Transaction) (id int, err error)
}

type TransactionInteractor struct {
}

func NewTransactionInteractor() TransactionInteractor {
	return &TransactionInteractor{}
}

func (interactor *TransactionInteractor) PostTransaction() (transactions []model.Transaction, err error) {
	transactions, _ = interactor.PostTransaction()
	return
}
