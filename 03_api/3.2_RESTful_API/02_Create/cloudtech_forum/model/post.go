package model

import "time"

// Postsテーブルに対応する構造体
type Post struct {
	ID        int       `json:"id"`         // ID
	Content   string    `json:"content"`    // 投稿内容
	UserID    int       `json:"user_id"`    // 投稿ユーザ
	CreatedAt time.Time `json:"created_at"` // 作成日
	UpdatedAt time.Time `json:"updated_at"` // 更新日
}
