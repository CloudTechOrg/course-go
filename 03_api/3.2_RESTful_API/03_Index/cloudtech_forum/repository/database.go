package repository

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// データベースの初期化を行う関数
func InitDB(user, password, host, port, dbname string) (err error) {
	// MySQLへの接続文字列を作成
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	// MySQLに接続
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("データベース接続エラー: %v", err)
		return err
	}

	// データベースの接続情報を返却
	return db.Ping()
}

// データベース接続を閉じる関数
func CloseDB() {
	if db != nil {
		// データベース接続を閉じる
		err := db.Close()
		if err != nil {
			log.Printf("データベースクローズエラー: %v", err)
		}
	}
}
