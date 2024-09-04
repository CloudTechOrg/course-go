package main

import (
	"fmt"
	"net/http"
)

// / パスに対応するハンドラ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "indexHandlerが実行されました。")
}

// /show パスに対応するハンドラ
func showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "showHandlerが実行されました。")
}

func main() {
	// ルーティングの設定
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/show", showHandler)

	// サーバの起動
	fmt.Println("HTTPサーバを起動しました。ポート: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTPサーバの起動に失敗しました: ", err)
	}
}
