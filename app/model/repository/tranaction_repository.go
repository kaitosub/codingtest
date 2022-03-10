package repository

import (
	"kaitosub/app/model/entity"
	"log"
)

// 外部パッケージに公開するインタフェース
type TransactionRepository interface {
	InsertTransaction(transaction entity.TransactionEntity) (id int, err error)
}

// 非公開のTodoRepository構造体
type transactionRepository struct {
}

// TodoRepositoryのコンストラクタ。TodoRepository構造体のポインタを返却する。
func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

// TODO追加処理
func (tr *transactionRepository) InsertTransaction(transaction entity.TransactionEntity) (id int, err error) {
	// 引数で受け取ったEntityの値を元にDBに追加
	_, err = DB.Exec("INSERT INTO transaction (user_id, amount, description) VALUES (?, ?, ?)", transaction.UserID, transaction.Amount, transaction.Description)
	if err != nil {
		log.Print(err)
		return
	}
	// created_atが最新のTODOのIDを返却
	err = DB.QueryRow("SELECT id FROM transaction ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}
