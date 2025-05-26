# ハンズオン手順
## 1. サインアップ処理の作成
1. `util/auth.go`ファイルを作成する
2. 下記のコードを記載する
    ```go
    // SECRET_HASH を計算する関数
    func calculateSecretHash(clientSecret, username, clientID string) string {
        mac := hmac.New(sha256.New, []byte(clientSecret))
        mac.Write([]byte(username + clientID))
        return base64.StdEncoding.EncodeToString(mac.Sum(nil))
    }
    ```

    ```go

    // Cognitoに新規ユーザーを登録する関数
    func Signup(
        clientID string,
        clientSecret string,
        email string,
        password string,
    ) (*cognitoidentityprovider.SignUpOutput, error) {

        // AWSセッションを作成
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String("ap-northeast-1"),
        }))
        svc := cognitoidentityprovider.New(sess)

        // クライアントシークレットを用いてSecretHashを計算
        secretHash := calculateSecretHash(clientSecret, email, clientID)

        // サインアップ用のリクエストを作成
        input := &cognitoidentityprovider.SignUpInput{
            ClientId:   aws.String(clientID),
            Username:   aws.String(email),
            Password:   aws.String(password),
            SecretHash: aws.String(secretHash),
            UserAttributes: []*cognitoidentityprovider.AttributeType{
                {
                    Name:  aws.String("email"),
                    Value: aws.String(email),
                },
            },
        }

        // Cognitoにサインアップリクエストを送信
        result, err := svc.SignUp(input)
        if err != nil {
            return nil, err
        }

        // 結果を返却
        return result, nil
    }

    ```

## 2. ハンドラーの作成
1. `handler/auth_handler.go`ファイルを作成する
2. 以下のコードを記載する
    ```go

    func SignupHandler(w http.ResponseWriter, r *http.Request) {
        // リクエストボディで受け取る項目を定義
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
            Name     string `json:"name"`
        }

        // リクエストボディの情報を読み取る
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
            return
        }

        // CognitoのクライアントIDとシークレットを環境変数から取得
        clientID := os.Getenv("COGNITO_CLIENT_ID")
        clientSecret := os.Getenv("COGNITO_CLIENT_SECRET")

        // サインアップ処理を実行
        result, err := auth.Signup(clientID, clientSecret, req.Email, req.Password)
        if err != nil {
            log.Println("Signupエラー:", err)
            http.Error(w, "サインアップに失敗しました", http.StatusInternalServerError)
            return
        }

        // ユーザー作成に成功した場合は201 Createdを返す
        w.WriteHeader(http.StatusCreated)

        // 成功メッセージを返す
        json.NewEncoder(w).Encode(map[string]string{
            "message": "ユーザー登録が完了しました",
            "userSub": *result.UserSub, // Cognitoが返すユーザーIDを含める（任意）
        })
    }

    ```

## 3. main.goの修正
1. `main.go`ファイルを開く
2. 下記内容を記載する
    ```go

    r.HandleFunc("/signup", handler.SignupHandler).Methods("POST")
    
    ```


## 4. 動作確認
1. ターミナルにて、下記のCURLコマンドを実行（`<your-mail-address>`と`<your-password>`はご自身のものに置き換え
    ```
    curl -X POST http://localhost:8080/signup \
    -H "Content-Type: application/json" \
    -d '{"email": "<your-mail-address>", "password": "<your-password>"}'
    ```
2. 指定したメールアドレスに認証コードが届くので、そのコードを控えておく