# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること

# ハンズオン手順
## 1. repositoryの更新
1. `repository/post_repository.go`ファイルを開き、下記内容を尽きする
    ```go 
    // 投稿の更新
    func UpdatePost(id int, content string, createdUserID int) (int64, error) {
      // SQLを定義
      query := "UPDATE posts SET content = ?, user_id = ? WHERE id = ?"

      // UPDATEのSQLを実行
      result, err := db.Exec(query, content, createdUserID, id)
      if err != nil {
        // エラーログの出力
        log.Printf("投稿の更新に失敗しました: %v", err)
        return 0, fmt.Errorf("投稿の更新に失敗しました: %w", err)
      }

      // 影響を受けた行数を取得
      rowsAffected, err := result.RowsAffected()
      if err != nil {
        // エラーログの出力
        log.Printf("影響を受けた行数の取得に失敗しました: %v", err)
        return 0, fmt.Errorf("影響を受けた行数の取得に失敗しました: %w", err)
      }

      // 影響を受けた行数を返す
      return rowsAffected, nil
    }
    ```

## 2. handlerの作成
1. `handler/post_handler.go`ファイルを開き、下記内容を追記する
    ```go
    // Updateハンドラ関数
    func Update(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
      id, _ := strconv.Atoi(vars["id"])

      // リクエストのBodyデータを格納するオブジェクトを定義
      var post model.Post
      if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
        return
      }

      // 更新処理の実行
      update_count, err := repository.UpdatePost(id, post.Content, post.UserID)
      if err != nil {
        log.SetFlags(log.LstdFlags | log.Lshortfile)
        http.Error(w, "更新処理に失敗しました", http.StatusInternalServerError)
        return
      }

      // 更新件数が0件の場合、404エラーを返す
      if update_count == 0 {
        http.Error(w, "更新対象のリソースが見つかりません", http.StatusNotFound)
        return
      }

      // レスポンスのBodyに更新件数をセット
      response := map[string]interface{}{
        "message":     "更新が成功しました",
        "updateCount": int(update_count),
      }

      // レスポンスを返却
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(response)
    }
    ```

## 3. main.goの作成
1. `main.go`ファイルを開き、下記内容を追記する
    ```go
    // ルーティングの設定
    r := mux.NewRouter()
    r.HandleFunc("/posts", handler.Create).Methods("POST")
    r.HandleFunc("/posts", handler.Index).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", handler.Show).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", handler.Update).Methods("PUT") // このコードを追加
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
1. 以下のcurlコマンドで、Updateの処理を実行する
  ```sh
  curl -X PUT http://localhost:8080/posts/2 \
    -H "Content-Type: application/json" \
    -d '{"content": "Go言語を覚えてきました", "user_id": 1}'
  ```

2. 以下のように、更新が成功したメッセージが表示されることを確認
  ```sh
  {"message":"更新が成功しました","updateCount":1}
  ```

3. 以下のコマンドで、IDが2のレコードを検索する
  ```sh
  curl -X GET http://localhost:8080/posts/2
  ```

4. `content`と`updated_at`が更新されていることを確認
  ```sh
  {"id":2,"content":"Go言語を覚えてきました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T23:05:04Z"}
  ```

## 6. APIサーバの停止
1. 起動している`APIサーバをControl + C`ボタンで停止する

## 7. GitHubへのPush
1. 以下のコマンドで、変更をコミットし、GitHubにプッシュする
    ```sh
    git add .
    git commit -m "add post update"
    git push origin feature/add-posts
    ```








# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

# ハンズオン手順

## HTTPサーバの起動

1. 下記のコマンドでフォルダーの移動を行う
```sh
cd 03_api/3.2_RESTful_API/05_Update/cloudtech_forum
```

2. Goのアプリケーションを実行し、HTTPサーバを起動する
```sh
$ go run main.go
```

3. 以下のような正常終了を示すメッセージが表示されることを確認
```sh
2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
```

## Updateの実行
1. 以下のcurlコマンドで、Updateの処理を実行する
```sh
curl -X PUT http://localhost:8080/posts/2 \
  -H "Content-Type: application/json" \
  -d '{"content": "Go言語を覚えてきました", "user_id": 1}'
```

2. 以下のように、更新が成功したメッセージが表示されることを確認
```sh
{"message":"更新が成功しました","updateCount":1}
```

3. 以下のコマンドで、IDが2のレコードを検索する
```sh
curl -X GET http://localhost:8080/posts/2
```

4. `content`と`updated_at`が更新されていることを確認
```sh
{"id":2,"content":"Go言語を覚えてきました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T23:05:04Z"}
```