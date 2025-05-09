# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること

# ハンズオン手順
## 1. ブランチの作成
1. 現時点のmainブランチを最新化する
    ```
    git pull
    ```

2. 以下のコマンドで、Gitの新しいブランチを作成する
    ```
    git checkout -b feature/add-posts
    ```

## 2. Goのモジュールの作成
1. 以下のコマンドで、`cloudtech-forum`という名称のGoのモジュールを作る
    ```
    go mod init cloudtech-forum
    ```
2. `go.mod`ファイルが作成されることを確認


## 3. modelの作成
1. `model`フォルダーを作成する
2. その中に`post.go`ファイルを作成し、下記内容を記載する
      ```go
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
      ```
## 4. repositoryの作成
1. `repository`フォルダーを作成する
2. その中に`database.go`ファイルを作成し、下記内容を記載する
    ```go
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
    ```
3. さらに`post_repository.go`ファイルを作成し、下記内容を記載する
    ```go 
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
    ```

## 5. handlerの作成
1. `handler`フォルダーを作成する
2. その中に`post_handler.go`ファイルを作成し、下記内容を記載する
    ```go
    package handler

    import (
      "encoding/json"
      "net/http"

      "cloudtech-forum/model"
      "cloudtech-forum/repository"
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
    ```
## 6. main.goの作成
1. `main.go`ファイルを作成し、下記内容を記載する
    ```go
    package main

    import (
      "log"
      "net/http"
      "os"

      "cloudtech-forum/handler"
      "cloudtech-forum/repository"

      "github.com/gorilla/mux"
      "github.com/joho/godotenv"
    )

    func init() {
      // .envファイルから環境変数を読み込む
      godotenv.Load()
    }

    func main() {

      // 環境変数からデータを取得
      apiport := os.Getenv("API_PORT")     // APIサーバのポート
      username := os.Getenv("DB_USERNAME") // DBのユーザ名
      password := os.Getenv("DB_PASSWORD") // DBのパスワード
      host := os.Getenv("DB_HOST")         // DBのホスト
      port := os.Getenv("DB_PORT")         // DBのポート
      dbname := os.Getenv("DB_NAME")       // DB名

      // データベースの接続を初期化
      err := repository.InitDB(username, password, host, port, dbname)
      if err != nil {
        log.Fatalf("データベースに接続できません: %v", err)
      }
      defer repository.CloseDB() // プログラム終了時にデータベース接続を閉じる

      // ルーティングの設定
      r := mux.NewRouter()
      r.HandleFunc("/posts", handler.Create).Methods("POST")

      // APIサーバを起動
      log.Println("APIサーバを起動しました。ポート: " + apiport)
      if err := http.ListenAndServe(":"+apiport, r); err != nil {
        log.Fatal(err)
      }
    }
    ```

## 7. packageの整理
3. 以下のコマンドで、必要なパッケージをインストールする
    ```go
    go mod tidy
    ```
4. `go.sum`が作成されることを確認


## 8. HTTPサーバの起動
1. Goのアプリケーションを実行し、HTTPサーバを起動する
    ```sh
    go run main.go
    ```

3. 以下のような正常終了を示すメッセージが表示されることを確認
    ```sh
    2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
    ```

## 8. 動作確認
1. 以下のcurlコマンドで、投稿データの登録を行う
    ```sh
    curl -X POST http://localhost:8080/posts \
      -H "Content-Type: application/json" \
      -d '{"content": "AWSはじめました", "user_id": 1}'
    ```

2. 下記のように正常終了を示すメッセージが表示されることを確認
    ```sh
    {"id":1,"message":"登録が成功しました"}
    ```

## 9. 登録データの確認

1. 以下のコマンドでMySQLにログイン
    ```
    mysql -u root -p
    ```

2. 以下のコマンドで、データの検索を行う
    ```sql
    select * from cloudtech_forum.posts;
    ```

3. 以下のように、登録したデータが表示されることを確認する
    ```
    +----+-----------------------+---------+---------------------+---------------------+
    | id | content               | user_id | created_at          | updated_at          |
    +----+-----------------------+---------+---------------------+---------------------+
    |  1 | AWSはじめました       |       1 | 2025-04-15 10:30:24 | 2025-04-15 10:30:24 |
    +----+-----------------------+---------+---------------------+---------------------+
    1 row in set (0.00 sec)
    ```

## 10. APIサーバの停止
1. 起動している`APIサーバをControl + C`ボタンで停止する

## 10. GitHubへのPush
1. 以下のコマンドで、変更をコミットし、GitHubにプッシュする
    ```sh
    git add .
    git commit -m "add post create"
    git push origin feature/add-posts
    ```