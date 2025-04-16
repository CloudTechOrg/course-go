# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

# ハンズオン手順

## HTTPサーバの起動

1. 下記のコマンドでフォルダーの移動を行う
```sh
cd 03_api/3.2_RESTful_API/03_Index/cloudtech_forum
```

2. Goのアプリケーションを実行し、HTTPサーバを起動する
```sh
$ go run main.go
```

3. 以下のような正常終了を示すメッセージが表示されることを確認
```sh
2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
```

## Indexの実行
1. 以下のcurlコマンドで、Indexの処理を実行する
```sh
curl -X GET http://localhost:8080/posts
```

2. 以下のように、ひとつ前に登録したデータが表示されること
```sh
[{"id":1,"content":"AWSはじめました","user_id":1,"created_at":"2025-04-15T10:30:24Z","updated_at":"2025-04-15T10:30:24Z"}]
```

## Indexの実行（複数件数）

1. 以下のコマンドで、テスト用のデータを追加で2件登録する

```sh
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"content": "Go言語はじめました", "user_id": 1}'
```

```sh
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"content": "Terraformはじめました", "user_id": 1}'
```

2. 以下のcurlコマンドで、Indexの処理を実行する
```sh
curl -X GET http://localhost:8080/posts
```

3. 以下のように、登録した3件のデータが表示されること

```sh
$ curl -X GET http://localhost:8080/posts
[{"id":1,"content":"AWSはじめました","user_id":1,"created_at":"2025-04-15T10:30:24Z","updated_at":"2025-04-15T10:30:24Z"},{"id":2,"content":"Go言語はじめました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T21:18:08Z"},{"id":3,"content":"Terraformはじめました","user_id":1,"created_at":"2025-04-16T21:18:13Z","updated_at":"2025-04-16T21:18:13Z"}]
```