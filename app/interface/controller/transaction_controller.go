package controller

import (
	"encoding/json"
	"github.com/kaitosub/codingtest/app/entity/model"
	"github.com/kaitosub/codingtest/app/infrastructure/mysql"
	"github.com/kaitosub/codingtest/app/interface/database"
	"github.com/kaitosub/codingtest/app/usecase"
	"log"
	"net/http"
	"strconv"
)

type TransactionController struct {
	interactor usecase.TransactionInteractorInterface
}

type TransactionControllerInterface interface {
	PostTransaction(w http.ResponseWriter, r *http.Request)
}

func NewTransactionController() TransactionControllerInterface {
	return &TransactionController{
		interactor: usecase.NewTransactionInteractor(&database.TransactionRepository{}),
	}
}

var amountLimit = 1000

func (controller *TransactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// リクエストbodyのJSONをDTOにマッピング
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var transactionRequest model.TransactionRequest
	json.Unmarshal(body, &transactionRequest)

	// DTOをTODOのEntityに変換
	transaction := model.TransactionEntity{
		UserID:      transactionRequest.UserId,
		Amount:      transactionRequest.Amount,
		Description: transactionRequest.Description,
	}

	// リポジトリの追加処理呼び出し
	id, err := InsertTransaction(transaction)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// LocationにリソースのPATHを設定し、ステータスコード２０１を返却
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
	//rr := http.Request{Method: "POST", URL: "http://localhost:8888/transactions", Response: nil, HTTP/1.1 1 1 map[Apikey:[secure-api-key-1] Content-Type:[application/json]] {} 0x1391e60 58 [] false localhost:8888 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc00001a110}

	//body := make([]byte, r.ContentLength)
	//_, err := r.Body.Read(body)
	//if err != nil {
	//	//err = string("err %s", hogehoge)
	//	//log.Fatalln("33:", err)
	//	return
	//}
	//var transactionRequest model.TransactionRequest
	//err = json.Unmarshal(body, &transactionRequest)
	//if err != nil {
	//	//log.Fatalln("39:", err)
	//	return
	//}
	//
	//transaction := model.TransactionEntity{UserID: transactionRequest.UserId, Amount: transactionRequest.Amount, Description: transactionRequest.Description}
	//log.Println(transaction)

	//var amount int
	//if err := mysql.DB.QueryRow(
	//	"select sum(amount) from transactions where user_id=?",
	//	ctxUser.ID,
	//).Scan(&amount); err != nil {
	//	log.Fatal(err)
	//}
	//if amount > amountLimit {
	//	log.Printf("amount %d over the amountLimit %d", amount, amountLimit)
	//}
	//
	//if amount+transaction.Amount > amountLimit {
	//	log.Print("error")
	//	return
	//}

	//ひつよう
	//id, err := InsertTransaction(transaction)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Fatalln("61:", err)
	//	return
	//}
	//
	//err = json.NewEncoder(w).Encode(transaction)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Fatalln("77:", err, r)
	//	return
	//}
	//
	//_, err = ioutil.ReadAll(r.Body)
	//if err != nil {
	//	return
	//}

	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(r.Body)
	//w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	//w.WriteHeader(201)
	//return
}

func InsertTransaction(transaction model.TransactionEntity) (id int, err error) {
	_, err = mysql.DB.Exec(
		"INSERT INTO transactions (user_id, amount, description) VALUES (?, ?, ?)",
		transaction.UserID, transaction.Amount, transaction.Description,
	)
	if err != nil {
		log.Println(err)
		return
	}
	// created_atが最新のTODOのIDを返却
	err = mysql.DB.QueryRow("SELECT id FROM transactions ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}
