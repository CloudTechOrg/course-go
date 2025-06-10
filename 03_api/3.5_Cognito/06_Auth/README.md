# ハンズオン手順
## 1. AccessTokenのチェック処理
1. `util/auth.go`ファイルを開く
2. 下記のコードを記載する
    ```go
    func CheckAccessToken(accessToken string) bool {
        jwksURL := "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_fr8xQq6eT/.well-known/jwks.json"

        options := keyfunc.Options{
            RefreshInterval: time.Hour,
            RefreshErrorHandler: func(err error) {
                log.Printf("JWKs refresh error: %v", err)
            },
        }

        jwks, err := keyfunc.Get(jwksURL, options)
        if err != nil {
            log.Printf("JWK取得失敗: %v", err)
            return false
        }

        token, err := jwt.Parse(accessToken, jwks.Keyfunc)
        if err != nil {
            log.Printf("JWT検証エラー: %v", err)
            return false
        }

        if !token.Valid {
            log.Println("トークンは無効です")
            return false
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            log.Println("クレームの取得に失敗しました")
            return false
        }

        log.Printf("スコープ: %v", claims["scope"])

        return true
    }

    // 認証チェック関数：成功すれば true、失敗すれば false を返す
    func IsAuthenticated(r *http.Request) bool {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            return false
        }

        const prefix = "Bearer "
        accessToken := strings.TrimPrefix(authHeader, prefix)

        return CheckAccessToken(accessToken)
    }
    ```

## 2. 各ハンドラー処理の冒頭にチェックを追加
1. `handler/post_handler.go`ファイルを開く
2. `Create`メソッドの冒頭に、下記を追加
    ```go
	if !auth.IsAuthenticated(r) {
		http.Error(w, "認証エラー", http.StatusUnauthorized)
		return
	}
    ```
3. `Index`メソッドの冒頭に、下記を追加
    ```go
	if !auth.IsAuthenticated(r) {
		http.Error(w, "認証エラー", http.StatusUnauthorized)
		return
	}
    ```
4. `Show`メソッドの冒頭に、下記を追加
    ```go
	if !auth.IsAuthenticated(r) {
		http.Error(w, "認証エラー", http.StatusUnauthorized)
		return
	}
    ```
5. `Update`メソッドの冒頭に、下記を追加
    ```go
	if !auth.IsAuthenticated(r) {
		http.Error(w, "認証エラー", http.StatusUnauthorized)
		return
	}
    ```
6. `Delete`メソッドの冒頭に、下記を追加
    ```go
	if !auth.IsAuthenticated(r) {
		http.Error(w, "認証エラー", http.StatusUnauthorized)
		return
	}
    ```

## 4. 動作確認（Access Token正常）
1. ターミナルにて、下記のCURLコマンドを実行（`<your-access-token>`はご自身のものに置き換え
    ```
    curl -X GET http://localhost:8080/posts \
    -H "Authorization: Bearer <your-access-token>"
    ```
2. 一覧が表示されることを確認

## 5. 動作確認（Access Tokenなし）
1. ターミナルにて、下記のCURLコマンドを実行（`<your-access-token>`はご自身のものに置き換え
    ```
    curl -X GET http://localhost:8080/posts
    ```
2. エラーとなることを確認
