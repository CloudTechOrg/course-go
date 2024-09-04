package main

import (
	"fmt"
	"net/http"
)

// GETメソッドのHandler
func indexHandler(w http.ResponseWriter) {
	fmt.Fprintln(w, "indexHandlerが実行されました")
}

// POSTメソッドのハンドラー
func createHandler(w http.ResponseWriter) {
	fmt.Fprintln(w, "createHandlerが実行されました")
}

// PUTメソッドのハンドラー
func updateHandler(w http.ResponseWriter) {
	fmt.Fprintln(w, "updateHandlerが実行されました")
}

// DELETEメソッドのハンドラー
func deleteHandler(w http.ResponseWriter) {
	fmt.Fprintln(w, "deleteHandlerが実行されました")
}

// ルーティングを処理する関数
func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		indexHandler(w)
	case http.MethodPost:
		createHandler(w)
	case http.MethodPut:
		updateHandler(w)
	case http.MethodDelete:
		deleteHandler(w)
	default:
		fmt.Fprintln(w, "対象外のHTTPメソッドが指定されました")
	}
}

func main() {
	// ルーティングの設定
	http.HandleFunc("/", routeHandler)

	// サーバの起動
	fmt.Println("HTTPサーバを起動しました。ポート: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTPサーバの起動に失敗しました: ", err)
	}
}
