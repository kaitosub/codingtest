package model

type Transaction struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type TransactionEntity struct {
	ID          int
	UserID      int
	Amount      int
	Description string
}

type TransactionRequest struct {
	UserId      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type TransactionAmountSum struct {
	Amount int `json:"amount"`
}

var TransactionQuery = transactionQuery{
	//FindTransactions:         qFindTransactions,
	FIndTransactionAmountSum: qFindTransactionAmountSum,
}

type transactionQuery struct {
	//FindTransactions         string
	FIndTransactionAmountSum string
}

//const qFindTransactions = `
//	SELECT
//		t.id,
//		t.user_id,
//		t.amount,
//		t.description
//	FROM transactions AS t
//	WHERE t.user_id = ?
//`

const qFindTransactionAmountSum = `
	SELECT
		sum(amount)
	FROM transactions
	WHERE user_id = ?
`
