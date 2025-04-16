package repository

import (
	"cloudtech_forum/model"
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

// 投稿の一覧検索
func SearchPostAll() ([]model.Post, error) {
	// SQLを定義
	query := "SELECT id, content, user_id, created_at, updated_at FROM posts"

	// SELECTのSQLを実行
	rows, err := db.Query(query)
	if err != nil {
		// エラーログの出力
		log.Printf("投稿の一覧検索に失敗しました: %v", err)
		return nil, fmt.Errorf("投稿の一覧検索に失敗しました: %w", err)
	}

	// rowsをclose
	defer rows.Close()

	// 一覧データを格納するスライスを定義
	var posts []model.Post

	// 一覧データを読み取りスライスに登録
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			// エラーログの出力
			log.Printf("投稿データの読み取りに失敗しました: %v", err)
			return nil, fmt.Errorf("投稿データの読み取りに失敗しました: %w", err)
		}
		posts = append(posts, post)
	}

	// 投稿データの一覧を返却
	return posts, nil
}
