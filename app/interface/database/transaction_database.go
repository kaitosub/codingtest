package database

import (
	"github.com/mfkessai/codetest-docker/app/entity/model"
	"github.com/mfkessai/codetest-docker/app/util/ctx"
)

var amountLimit = 1000

type TransactionRepository struct{}

type TransactionRepositoryInterface interface {
	FindTransactions(*ctx.CtxUser) []model.Transaction
}

func (tr *TransactionRepository) FindTransactions(ctxUser *ctx.CtxUser) []model.Transaction {

	var transactions []model.Transaction
	//var transactionAmountSum []model.TransactionAmountSum
	//q := model.TransactionQuery.FindTransactions

	//conn, err := sql.Open("mysql", "root@tcp(127.0.0.1)/codetest")
	//if err != nil {
	//	log.Println(err)
	//}

	//var amount int
	//if err := conn.QueryRow(
	//	"select sum(amount) from transactions where user_id=?",
	//	ctxUser.ID,
	//).Scan(&amount); err != nil {
	//	log.Println(err)
	//}
	//if amount > amountLimit {
	//	log.Println("amount %d over the amountLimit %d", amount, amountLimit)
	//}

	//TransactionAmount := model.TransactionQuery.FindTransactionAmountSum
	//err := mysql.DB.Select(&transactionAmountSum, TransactionAmount, ctxUser.ID)
	//if err != nil {
	//	return transactionAmountSum, err
	//}

	//err := mysql.DB.Select(&transactions, q, ctxUser.ID)
	//if err != nil {
	//	//logger.ErrorStack(err)
	//	return transactions, err
	//}

	return transactions
}
