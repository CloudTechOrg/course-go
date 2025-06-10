# ハンズオン手順
## 1. サインアップ処理の作成
1. `util/auth.go`ファイルを開く
2. 下記のコードを記載する
    ```go
    func Logout(accessToken string) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String("ap-northeast-1"),
        }))
        svc := cognitoidentityprovider.New(sess)

        input := &cognitoidentityprovider.GlobalSignOutInput{
            AccessToken: aws.String(accessToken),
        }

        result, err := svc.GlobalSignOut(input)
        if err != nil {
            return nil, err
        }
        return result, nil
    }
    ```

## 2. ハンドラーの作成
1. `handler/auth_handler.go`ファイルを開く
2. 以下のコードを記載する
    ```go
    // Logoutハンドラ関数
    func LogoutHandler(w http.ResponseWriter, r *http.Request) {
        // Authorizationヘッダーからアクセストークンを取得
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "アクセストークンが指定されていません", http.StatusUnauthorized)
            return
        }

        // "Bearer " プレフィックスを除去（必要に応じて）
        accessToken := authHeader
        if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
            accessToken = authHeader[7:]
        }

        _, err := auth.Logout(accessToken)
        if err != nil {
            http.Error(w, "サインアウトに失敗しました", http.StatusInternalServerError)
            return
        }

        response := map[string]string{
            "message": "サインアウトに成功しました",
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }
    ```

## 3. main.goの修正
1. `main.go`ファイルを開く
2. 下記内容を記載する
    ```go
	r.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")
    ```

## 4. 動作確認
1. 下記のCURLコマンドを実行（'<your-access-token>'はご自身の内容に置き換え）
    ```
    curl -X POST http://localhost:8080/logout \
    -H "Authorization: Bearer <あなたのアクセストークン>" \
    -H "Content-Type: application/json"
    ```
2. 正常終了することを確認