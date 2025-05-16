# ハンズオン手順
## 1. サインアップ処理の作成
1. `util/auth.go`ファイルを開く
2. 下記のコードを記載する
    ```go
    func ConfirmCode(clientID string, clientSecret string, email string, confirmationCode string) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
        // AWSセッションの初期化
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String("ap-northeast-1"),
        }))
        svc := cognitoidentityprovider.New(sess)

        // SECRET_HASH の計算
        secretHash := calculateSecretHash(clientSecret, email, clientID)

        input := &cognitoidentityprovider.ConfirmSignUpInput{
            ClientId:         aws.String(clientID),
            Username:         aws.String(email),
            ConfirmationCode: aws.String(confirmationCode),
            SecretHash:       aws.String(secretHash),
        }

        result, err := svc.ConfirmSignUp(input)
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
    // ConfirmSignupハンドラ関数
    func ConfirmSignupHandler(w http.ResponseWriter, r *http.Request) {
        var req struct {
            Email            string `json:"email"`
            ConfirmationCode string `json:"confirmation_code"`
        }

        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
            return
        }

        clientID := os.Getenv("COGNITO_CLIENT_ID")
        clientSecret := os.Getenv("COGNITO_CLIENT_SECRET")

        _, err := auth.ConfirmCode(clientID, clientSecret, req.Email, req.ConfirmationCode)
        if err != nil {
            http.Error(w, "確認コードの検証に失敗しました", http.StatusBadRequest)
            return
        }

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