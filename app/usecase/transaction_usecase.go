package usecase

import (
	"github.com/kaitosub/codingtest/app/entity/model"
	"github.com/kaitosub/codingtest/app/interface/database"
)

type TransactionInteractor struct {
	repository database.TransactionRepositoryInterface
}

func NewTransactionInteractor(t database.TransactionRepositoryInterface) TransactionInteractorInterface {
	return &TransactionInteractor{
		repository: t,
	}
}

type TransactionInteractorInterface interface {
	PostTransaction(transaction model.Transaction) (t model.Transaction, err error)
}

func (interactor *TransactionInteractor) PostTransaction(transaction model.Transaction) (transactions model.Transaction, err error) {
	transactions, _ = interactor.PostTransaction(transaction)
	return
}
