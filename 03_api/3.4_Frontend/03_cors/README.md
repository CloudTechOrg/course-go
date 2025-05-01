# ハンズオン手順

## 1. CORSの実装
1. ローカルにて、cloudtech-forumのリポジトリを開く

2. `main.go`を開く

3. CORS対策を行う`enableCORS`関数を、`main.go`に追加する
    ```go
    // CORSミドルウェア
    func enableCORS(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // CORSヘッダーを設定（必要に応じて制限することも可能）
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

            // OPTIONSメソッドへの即時レスポンス（Preflightリクエスト対応）
            if r.Method == http.MethodOptions {
                w.WriteHeader(http.StatusOK)
                return
            }

            // 通常処理
            next.ServeHTTP(w, r)
        })
    }
    ```

4. 上記のenableCORS関数を、main関数から呼び出す
    ```go
    func main() {

        // 環境変数からデータを取得
        apiport := os.Getenv("API_PORT")
        username := os.Getenv("DB_USERNAME")
        password := os.Getenv("DB_PASSWORD")
        host := os.Getenv("DB_HOST")
        port := os.Getenv("DB_PORT")
        dbname := os.Getenv("DB_NAME")

        // データベース接続
        err := repository.InitDB(username, password, host, port, dbname)
        if err != nil {
            log.Fatalf("データベースに接続できません: %v", err)
        }
        defer repository.CloseDB()

        // ルーター定義
        r := mux.NewRouter()
        r.HandleFunc("/posts", handler.Create).Methods("POST")
        r.HandleFunc("/posts", handler.Index).Methods("GET")
        r.HandleFunc("/posts/{id:[0-9]+}", handler.Show).Methods("GET")
        r.HandleFunc("/posts/{id:[0-9]+}", handler.Update).Methods("PUT")
        r.HandleFunc("/posts/{id:[0-9]+}", handler.Delete).Methods("DELETE")

        // CORSミドルウェアを適用
        corsRouter := enableCORS(r) // これを追加

        // APIサーバ起動
        log.Println("APIサーバを起動しました。ポート: " + apiport)
        if err := http.ListenAndServe(":"+apiport, corsRouter); err != nil {
            log.Fatal(err)
        }
    }
    ```

5. 以下のコマンドで、APIサーバを起動する
    ```shell
    go run main.go
    ```

6. フロント側のindex.htmlを実行し、一覧にデータが表示されることを確認する

7. 変更をGitHubに反映させる
    ```shell
    git add .
    git commit -m "add cors"
    git push origin main
    ```

## 2. CORSの変更をEC2インスタンスに反映
1. APIサーバのEC2インスタンスにsshなどでログインする

2. GitHubの変更内容をダウンロード（`git pull`）する
    ```shell
    cd /home/ec2-user/cloudtech-forum
    git pull origin main
    ```

3. サービスを再起動して変更を反映
    ```
    sudo systemctl restart goserver.service
    ```