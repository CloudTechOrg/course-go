package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSONで受け取るデータを構造体として定義
type RequestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	// リクエストボディを読み取る
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bodyの読み取りに失敗しました", http.StatusBadRequest)
		return
	}

	// JSONデータを構造体にデコードする
	var data RequestData
	err = json.Unmarshal(body, &data)
	if err != nil {
		// JSONのデコードに失敗した場合はエラーとする
		http.Error(w, "JSONのデコードに失敗しました", http.StatusBadRequest)
		return
	}

	// データの使用例
	fmt.Fprintf(w, "あなたが入力したデータは、名前: %s, 年齢: %d です。", data.Name, data.Age)
}

func main() {
	// ルーティングの設定
	http.HandleFunc("/", handler)

	// サーバの起動
	fmt.Println("HTTPサーバを起動しました。ポート: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTPサーバの起動に失敗しました: ", err)
	}
}
