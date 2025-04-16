# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

# ハンズオン手順

## HTTPサーバの起動

1. 下記のコマンドでフォルダーの移動を行う
```sh
cd 03_api/3.2_RESTful_API/05_Delete/cloudtech_forum
```

2. Goのアプリケーションを実行し、HTTPサーバを起動する
```sh
$ go run main.go
```

3. 以下のような正常終了を示すメッセージが表示されることを確認
```sh
2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
```

## Deleteの実行
1. 以下のcurlコマンドで、Deleteの処理を実行する
```sh
curl -X DELETE http://localhost:8080/posts/2
```

2. 以下のように、削除が成功したメッセージが表示されることを確認
```sh
{"deletedCount":1,"message":"削除が成功しました"}
```

3. 以下のコマンドで、postsテーブルの一覧を表示する
```sh
curl -X GET http://localhost:8080/posts
```

4. `id` が`2`の投稿データが削除されていることを確認する
```sh
[{"id":1,"content":"AWSはじめました","user_id":1,"created_at":"2025-04-15T10:30:24Z","updated_at":"2025-04-15T10:30:24Z"},{"id":3,"content":"Terraformはじめました","user_id":1,"created_at":"2025-04-16T21:18:13Z","updated_at":"2025-04-16T21:18:13Z"}]
```

5. 以下のように、存在しない`id`を削除しようとしてみる
```sh
curl -X DELETE http://localhost:8080/posts/4
```

6. エラーメッセージが表示されることを確認する
```sh
削除対象のリソースが見つかりません
```