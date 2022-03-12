package dto

type TransactionResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type TransactionRequest struct {
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"todos"`
}
