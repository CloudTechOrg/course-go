package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// リクエストデータの構造体を定義
type RequestData struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

// レスポンスデータの構造体を定義
type ResponseData struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Age      int    `json:"age"`
}

// 年齢を計算する関数
func calculateAge(birthday string) (int, error) {
	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, birthday)
	if err != nil {
		return 0, err
	}
	today := time.Now()
	age := today.Year() - birthDate.Year()
	if today.YearDay() < birthDate.YearDay() {
		age--
	}
	return age, nil
}

func handler(response http.ResponseWriter, request *http.Request) {
	// リクエストボディを読み取る
	request_json, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "bodyの読み取りに失敗しました", http.StatusInternalServerError)
		return
	}

	// JSONデータを構造体にデコードする
	var request_data RequestData
	err = json.Unmarshal(request_json, &request_data)
	if err != nil {
		http.Error(response, "JSONの形式が誤っています", http.StatusBadRequest)
		return
	}

	// 年齢を計算する
	age, err := calculateAge(request_data.Birthday)
	if err != nil {
		http.Error(response, "生年月日の形式が正しくありません。", http.StatusBadRequest)
		return
	}

	// レスポンスデータを作成
	response_data := ResponseData{
		Name:     request_data.Name,
		Birthday: request_data.Birthday,
		Age:      age,
	}

	// Content-Typeヘッダーにapplication/jsonを設定する
	response.Header().Set("Content-Type", "application/json")

	// 構造体をJSON形式に変換し、レスポンスのbodyに設定
	json.NewEncoder(response).Encode(response_data)
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
