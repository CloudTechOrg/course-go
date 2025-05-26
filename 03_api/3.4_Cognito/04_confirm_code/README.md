# ハンズオン手順
## 1. サインアップ処理の作成
1. `util/auth.go`ファイルを開く
2. 下記のコードを記載する
    ```go

    // メールに送信された確認コードを使って、Cognitoでユーザーを有効化する関数
    func ConfirmCode(
        clientID string,
        clientSecret string,
        email string,
        confirmationCode string,
    ) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
        // AWSセッションを初期化（リージョンは東京）
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String("ap-northeast-1"),
        }))

        // Cognitoクライアントを作成
        svc := cognitoidentityprovider.New(sess)

        // シークレットハッシュを計算
        secretHash := calculateSecretHash(clientSecret, email, clientID)

        // 確認コードとユーザー情報を設定
        input := &cognitoidentityprovider.ConfirmSignUpInput{
            ClientId:         aws.String(clientID),
            Username:         aws.String(email),
            ConfirmationCode: aws.String(confirmationCode),
            SecretHash:       aws.String(secretHash),
        }

        // サインアップ確認を実行
        result, err := svc.ConfirmSignUp(input)
        if err != nil {
            return nil, err
        }

        // 結果を返す
        return result, nil
    }

    ```

## 2. ハンドラーの作成
1. `handler/auth_handler.go`ファイルを開く
2. 以下のコードを記載する
    ```go

    // ConfirmSignupハンドラ関数
    func ConfirmSignupHandler(w http.ResponseWriter, r *http.Request) {
        // リクエストボディの構造体を定義（メールアドレスと確認コードを受け取る）
        var req struct {
            Email            string `json:"email"`
            ConfirmationCode string `json:"confirmation_code"`
        }

        // JSONデコード処理（失敗した場合は400エラーを返す）
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
            return
        }

        // 環境変数からCognitoのクライアントIDとシークレットを取得
        clientID := os.Getenv("COGNITO_CLIENT_ID")
        clientSecret := os.Getenv("COGNITO_CLIENT_SECRET")

        // 確認コードをCognitoに送信してサインアップを確定
        _, err := auth.ConfirmCode(clientID, clientSecret, req.Email, req.ConfirmationCode)
        if err != nil {
            http.Error(w, "確認コードの検証に失敗しました", http.StatusBadRequest)
            return
        }

        // 成功時のレスポンスをJSON形式で返す
        response := map[string]string{
            "message": "サインアップの確認が完了しました",
        }
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }

    ```

## 3. main.goの修正
1. `main.go`ファイルを開く
2. 下記内容を記載する
    ```go

	r.HandleFunc("/confirmcode", handler.ConfirmSignupHandler).Methods("POST")
    
    ```

## 4. 動作確認
1. 以下のCURLコマンドを実行する（`<your-mail-address>`はご自身のものに、`<your-confirm-code>`はサインアップ時にメールで受信したコードを入力）
    ```
    curl -X POST http://localhost:8080/confirmcode \
    -H "Content-Type: application/json" \
    -d '{"email": "<your-mail-address>", "confirmation_code": "<your-confirm-code>"}
    ```
2. 正常終了することを確認する