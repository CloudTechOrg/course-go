# ハンズオン手順
## 1. サインアップ処理の作成
1. `util/auth.go`ファイルを開く
2. 下記のコードを記載する
    ```go
    func Login(clientID string, clientSecret string, email string, password string) (*cognitoidentityprovider.AuthenticationResultType, error) {
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String("ap-northeast-1"),
        }))
        svc := cognitoidentityprovider.New(sess)

        secretHash := calculateSecretHash(clientSecret, email, clientID)

        input := &cognitoidentityprovider.InitiateAuthInput{
            AuthFlow: aws.String("USER_PASSWORD_AUTH"),
            ClientId: aws.String(clientID),
            AuthParameters: map[string]*string{
                "USERNAME":    aws.String(email),
                "PASSWORD":    aws.String(password),
                "SECRET_HASH": aws.String(secretHash),
            },
        }

        resp, err := svc.InitiateAuth(input)
        if err != nil {
            return nil, err
        }

        return resp.AuthenticationResult, nil
    }
    ```

## 2. ハンドラーの作成
1. `handler/auth_handler.go`ファイルを開く
2. 以下のコードを記載する
    ```go
    // Loginハンドラ関数
    func LoginHandler(w http.ResponseWriter, r *http.Request) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
            return
        }

        clientID := os.Getenv("COGNITO_CLIENT_ID")
        clientSecret := os.Getenv("COGNITO_CLIENT_SECRET")

        authResult, err := auth.Login(clientID, clientSecret, req.Email, req.Password)
        if err != nil {
            http.Error(w, "ログインに失敗しました", http.StatusUnauthorized)
            return
        }

        response := map[string]string{
            "id_token":      aws.StringValue(authResult.IdToken),
            "access_token":  aws.StringValue(authResult.AccessToken),
            "refresh_token": aws.StringValue(authResult.RefreshToken),
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }
    ```

## 3. main.goの修正
1. `main.go`ファイルを開く
2. 下記内容を記載する
    ```go
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")
    ```

## 4. 動作確認
1. ターミナルにて、下記のCURLコマンドを実行（`<your-mail-address>`と`<your-password>`はご自身のものに置き換え
    ```
    curl -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d '{"email": "kymx1983@gmail.com", "password": "ExamplePassword123!"}'
    ```
2. Access Tokenが返却されるので、控えておく