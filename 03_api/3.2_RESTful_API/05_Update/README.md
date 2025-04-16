# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

# ハンズオン手順

## HTTPサーバの起動

1. 下記のコマンドでフォルダーの移動を行う
```sh
cd 03_api/3.2_RESTful_API/04_Show/cloudtech_forum
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