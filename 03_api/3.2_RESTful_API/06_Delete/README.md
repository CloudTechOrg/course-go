# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること

# ハンズオン手順
## 1. repositoryの更新
1. `repository/post_repository.go`ファイルを開き、下記内容を尽きする
    ```go 
    // 投稿の削除
    func DeletePost(id int) (int64, error) {
        // SQLを定義
        query := "DELETE FROM posts WHERE id = ?"

        // DELETEのSQLを実行
        result, err := db.Exec(query, id)
        if err != nil {
            // エラーログの出力
            log.Printf("投稿の削除に失敗しました: %v", err)
            return 0, fmt.Errorf("投稿の削除に失敗しました: %w", err)
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
    // Deleteハンドラ関数
    func DeleteHandler(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, _ := strconv.Atoi(vars["id"])

        // 削除処理の実行
        delete_count, err := repository.DeletePost(id)
        if err != nil {
            http.Error(w, "削除処理に失敗しました", http.StatusInternalServerError)
            return
        }

        // 削除件数が0件の場合、404エラーを返す
        if delete_count == 0 {
            http.Error(w, "削除対象のリソースが見つかりません", http.StatusNotFound)
            return
        }

        // レスポンスのBodyに削除件数をセット
        response := map[string]interface{}{
            "message":      "削除が成功しました",
            "deletedCount": int(delete_count),
        }

        // ステータスコード200を返す
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
    r.HandleFunc("/posts/{id:[0-9]+}", handler.Update).Methods("PUT")
    r.HandleFunc("/posts/{id:[0-9]+}", handler.Delete).Methods("DELETE") // このコードを追加
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
1. 以下のcurlコマンドで、Deleteの処理を実行する
    ```sh
    curl -X DELETE http://localhost:8080/posts/2
    ```

2. 以下のように、削除が成功したメッセージが表示されることを確認
    ```sh
    {"deletedCount":1,"message":"削除が成功しました"}
    ```

3. 以下のコマンドで、postsテーブルの一覧を表示する
    ```sh
    curl -X GET http://localhost:8080/posts
    ```

4. `id` が`2`の投稿データが削除されていることを確認する
    ```sh
    [{"id":1,"content":"AWSはじめました","user_id":1,"created_at":"2025-04-15T10:30:24Z","updated_at":"2025-04-15T10:30:24Z"},{"id":3,"content":"Terraformはじめました","user_id":1,"created_at":"2025-04-16T21:18:13Z","updated_at":"2025-04-16T21:18:13Z"}]
    ```

5. 以下のように、存在しない`id`を削除しようとしてみる
    ```sh
    curl -X DELETE http://localhost:8080/posts/4
    ```

6. エラーメッセージが表示されることを確認する
    ```sh
    削除対象のリソースが見つかりません
    ```

## 6. APIサーバの停止
1. 起動している`APIサーバをControl + C`ボタンで停止する

## 7. GitHubへのPush
1. 以下のコマンドで、変更をコミットし、GitHubにプッシュする
    ```sh
    git add .
    git commit -m "add post delete"
    git push origin feature/add-posts
    ```

## 7. mainブランチにマージ
1. GitHubのリポジトリ画面を開く
2. 上部の「Pull requests」タブをクリック
3. 「New pull request」をクリック
4. baseをmain、compareをfeature/add-postsに設定
5. 差分を確認し、「Create pull request」をクリック
6. 必要に応じてタイトルやコメントを記入し、もう一度「Create pull request」をクリック
7. 問題なければ「Merge pull request」をクリックしてマージ完了

