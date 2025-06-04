# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること

# ハンズオン手順
## 1. repositoryの更新
1. `repository/post_repository.go`ファイルを開き、下記内容を尽きする
    ```go 
    // 投稿の詳細検索
    func SearchPost(id int) (model.Post, error) {
        // 検索結果を格納する構造体を定義
        var post model.Post

        // SQLを定義
        query := "SELECT id, content, user_id, created_at, updated_at FROM posts WHERE id = ?"

        // SELECTのSQLを実行
        row := db.QueryRow(query, id)

        // 読み込んだデータを、postの各フィールドに格納
        err := row.Scan(&post.ID, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            // エラーログの出力
            log.Printf("投稿の詳細検索に失敗しました: %v", err)
            return post, fmt.Errorf("投稿の詳細検索に失敗しました: %w", err)
        }

        // 読み込んだ投稿データを返却
        return post, nil
    }
    ```

## 2. handlerの作成
1. `handler/post_handler.go`ファイルを開き、下記内容を追記する
    ```go
    // Showハンドラ関数
    func ShowHandler(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, _ := strconv.Atoi(vars["id"])

        // 検索処理の実行
        post, err := repository.SearchPost(id)
        if err != nil {
            log.SetFlags(log.LstdFlags | log.Lshortfile)
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        // ステータスコードに「200：OK」を設定
        w.WriteHeader(http.StatusOK)

        // レスポンスのBodyに、検索した投稿データを設定
        json.NewEncoder(w).Encode(post)
    }
    ```

## 3. main.goの作成
1. `main.go`ファイルを開き、下記内容を追記する
    ```go
    // ルーティングの設定
    r := mux.NewRouter()
	r.HandleFunc("/posts", handler.Create).Methods("POST")
	r.HandleFunc("/posts", handler.Index).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handler.ShowHandler).Methods("GET")// このコードを追加
    ```

## 4. HTTPサーバの起動
1. Goのアプリケーションを実行し、HTTPサーバを起動する
    ```sh
    go run main.go
    ```

2. 以下のような正常終了を示すメッセージが表示されることを確認
    ```sh
    2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
    ```

## 5. 動作確認
1. 以下のcurlコマンドで、Showの処理を実行する
    ```sh
    curl -X GET http://localhost:8080/posts/2
    ```

2. 以下のように、idが`2`のデータが表示されること
    ```sh
    {"id":2,"content":"Go言語はじめました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T21:18:08Z"}
    ```

## 6. 動作確認（データが存在しない場合）
1. 以下のcurlコマンドで、Showの処理を実行する
    ```sh
    curl -X GET http://localhost:8080/posts/4
    ```

2. 以下のように、エラーメッセージが表示されること
    ```sh
    投稿の詳細検索に失敗しました: sql: no rows in result set
    ```

## 6. APIサーバの停止
1. 起動している`APIサーバをControl + C`ボタンで停止する

## 7. GitHubへのPush
1. 以下のコマンドで、変更をコミットし、GitHubにプッシュする
    ```sh
    git add .
    git commit -m "add post show"
    git push origin feature/add-posts
    ```