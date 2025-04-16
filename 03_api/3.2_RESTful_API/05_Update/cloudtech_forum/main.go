package main

import (
	"log"
	"net/http"
	"os"

	"cloudtech_forum/handler"
	"cloudtech_forum/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// .envファイルから環境変数を読み込む
	godotenv.Load()
}

func main() {

	// 環境変数からデータを取得
	apiport := os.Getenv("API_PORT")     // APIサーバのポート
	username := os.Getenv("DB_USERNAME") // DBのユーザ名
	password := os.Getenv("DB_PASSWORD") // DBのパスワード
	host := os.Getenv("DB_HOST")         // DBのホスト
	port := os.Getenv("DB_PORT")         // DBのポート
	dbname := os.Getenv("DB_NAME")       // DB名

	// データベースの接続を初期化
	err := repository.InitDB(username, password, host, port, dbname)
	if err != nil {
		log.Fatalf("データベースに接続できません: %v", err)
	}
	defer repository.CloseDB() // プログラム終了時にデータベース接続を閉じる

	// ルーティングの設定
	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.Create).Methods("POST")
	r.HandleFunc("/posts", handler.Index).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handler.Show).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handler.Update).Methods("PUT")

	// APIサーバを起動
	log.Println("APIサーバを起動しました。ポート: " + apiport)
	if err := http.ListenAndServe(":"+apiport, r); err != nil {
		log.Fatal(err)
	}
}
