package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// hello worldという文字列をレスポンスとして返す
	fmt.Fprintf(w, "hello world")
}

func main() {
	// /パスにアクセスがあった場合に、helloHandler関数を実行するように設定
	http.HandleFunc("/", helloHandler)

	// 8080ポートでサーバーを起動
	fmt.Println("HTTPサーバを起動しました。ポート: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTPサーバの起動に失敗しました: ", err)
	}
}
