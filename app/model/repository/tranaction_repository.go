package repository

import (
	"github.com/kaitosub/codingtest/app/model/entity"
	"log"
)

// 外部パッケージに公開するインタフェース
type TransactionRepository interface {
	InsertTransaction(transaction entity.TransactionEntity) (err error)
}

// 非公開のTodoRepository構造体
type transactionRepository struct {
}

// TodoRepositoryのコンストラクタ。TodoRepository構造体のポインタを返却する。
func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

// TODO追加処理
func (tr *transactionRepository) InsertTransaction(transaction entity.TransactionEntity) (err error) {
	// 引数で受け取ったEntityの値を元にDBに追加
	_, err = DB.Exec("INSERT INTO transactions (user_id, amount, description) VALUES (?, ?, ?)", transaction.UserID, transaction.Amount, transaction.Description)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
