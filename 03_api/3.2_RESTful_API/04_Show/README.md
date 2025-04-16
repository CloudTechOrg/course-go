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

## Showの実行
1. 以下のcurlコマンドで、Showの処理を実行する
```sh
curl -X GET http://localhost:8080/posts/2
```

2. 以下のように、idが`2`のデータが表示されること
```sh
{"id":2,"content":"Go言語はじめました","user_id":1,"created_at":"2025-04-16T21:18:08Z","updated_at":"2025-04-16T21:18:08Z"}
```

## Showの実行（データが存在しない場合）

```sh
curl -X GET http://localhost:8080/posts/4
```

2. 以下のように、エラーメッセージが表示されること
```sh
投稿の詳細検索に失敗しました: sql: no rows in result set
```