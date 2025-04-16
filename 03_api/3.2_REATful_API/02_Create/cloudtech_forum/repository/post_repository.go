package repository

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 投稿の新規登録
func CreatePost(content string, createdUserID int) (int, error) {
	// SQLを定義
	query := "INSERT INTO posts (content, user_id) VALUES (?, ?)"

	// INSERTのSQLを実行
	result, err := db.Exec(query, content, createdUserID)
	if err != nil {
		// エラーログの出力
		log.Printf("投稿の新規登録に失敗しました: %v", err)
		return 0, fmt.Errorf("投稿の新規登録に失敗しました: %w", err)
	}

	// 新規登録されたレコードのIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		// エラーログの出力
		log.Printf("新規登録IDの取得に失敗しました: %v", err)
		return 0, fmt.Errorf("新規登録IDの取得に失敗しました: %w", err)
	}

	// IDを返却
	return int(id), nil
}
