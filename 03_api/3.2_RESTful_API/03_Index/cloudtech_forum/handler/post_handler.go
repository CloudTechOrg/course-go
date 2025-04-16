package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"cloudtech_forum/model"
	"cloudtech_forum/repository"
)

// Createハンドラ関数
func Create(w http.ResponseWriter, r *http.Request) {
	// リクエストのBodyデータを格納するオブジェクトを定義
	var post model.Post

	// リクエストのBodyからJSONデータを取得し、post構造体に格納
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
		return
	}

	// 登録処理を実行
	id, err := repository.CreatePost(post.Content, post.UserID)
	if err != nil {
		http.Error(w, "投稿データの登録に失敗しました", http.StatusInternalServerError)
		return
	}

	// レスポンスのBodyに、登録されたレコードのIDを指定
	response := map[string]interface{}{
		"message": "登録が成功しました",
		"id":      id,
	}

	// レスポンスを返却
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Indexハンドラ関数
func Index(w http.ResponseWriter, r *http.Request) {
	// 検索処理の実行
	posts, err := repository.SearchPostAll()
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Printf("Error: %v", err)
		return
	}

	// ステータスコードに「200：OK」を設定
	w.WriteHeader(http.StatusOK)

	// postsデータのスライスをレスポンスとして設定
	json.NewEncoder(w).Encode(posts)
}
