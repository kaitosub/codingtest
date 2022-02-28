package usecase

import (
	"context"
	"github.com/mfkessai/codetest-docker/app/entity/model"
	"github.com/mfkessai/codetest-docker/app/interface/database"
	"github.com/mfkessai/codetest-docker/app/util/ctx"
)

type TransactionInteractor struct {
	repository database.TransactionRepositoryInterface
}

func NewTransactionInteractor(tr database.TransactionRepositoryInterface) TransactionInteractorInterface {
	return &TransactionInteractor{repository: tr}
}

type TransactionInteractorInterface interface {
	FindTransactions(context.Context) ([]model.Transaction, error)
}

func (interactor *TransactionInteractor) FindTransactions(
	context context.Context) (transactions []model.Transaction, err error) {
	ctxUser, err := ctx.GetCtxUser(context)
	if err != nil {
		return
	}

	transactions = interactor.repository.FindTransactions(ctxUser)
	return
}
