# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること
- GitHubアカウントが作成されていること

# ハンズオン手順
## 1. repositoryの更新
1. `repository/post_repository.go`ファイルを開き、下記内容を尽きする
    ```go 
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
    ```

## 2. handlerの作成
1. `handler/post_handler.go`ファイルを開き、下記内容を追記する
    ```go
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
    ```

## 3. main.goの作成
1. `main.go`ファイルを開き、下記内容を追記する
    ```go
    // ルーティングの設定
    r := mux.NewRouter()
    r.HandleFunc("/posts", handler.Create).Methods("POST")
    r.HandleFunc("/posts", handler.Index).Methods("GET")  // このコードを追加
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
1. 以下のコマンドで、テスト用のデータを追加で2件登録する

    ```sh
    curl -X POST http://localhost:8080/posts \
      -H "Content-Type: application/json" \
      -d '{"content": "Go言語はじめました", "user_id": 1}'
    ```

    ```sh
    curl -X POST http://localhost:8080/posts \
      -H "Content-Type: application/json" \
      -d '{"content": "Terraformはじめました", "user_id": 1}'
    ```

2. 以下のcurlコマンドで、Indexの処理を実行する
    ```sh
    curl -X GET http://localhost:8080/posts
    ```

3. 以下のように、登録した3件のデータが表示されること
    ```sh
    $ curl -X GET http://localhost:8080/posts
    [{"id":1,"content":"AWSはじめました","user_id":1,"created_at":"2025-04-15T10:30:24Z","updated_at":"2025-04-15T10:30:24Z"},{"id":2,"content":"Go言語はじめました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T21:18:08Z"},{"id":3,"content":"Terraformはじめました","user_id":1,"created_at":"2025-04-16T21:18:13Z","updated_at":"2025-04-16T21:18:13Z"}]
    ```

## 6. APIサーバの停止
1. 起動している`APIサーバをControl + C`ボタンで停止する

## 7. GitHubへのPush
1. 以下のコマンドで、変更をコミットし、GitHubにプッシュする
    ```sh
    git add .
    git commit -m "add post index"
    git push origin feature/add-posts
    ```