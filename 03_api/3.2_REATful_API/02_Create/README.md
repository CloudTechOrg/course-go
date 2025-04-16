# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

# ハンズオン手順

## MySQLの初期データ登録
1. 以下のコマンドでMySQLにログイン
```
mysql -u root -p
```

2. 以下のコマンドで、`cloudtech_forum`のデータベースを作成
```sql
CREATE DATABASE cloudtech_forum;
```

3. 以下のコマンドで、postsのテーブルを作成する
```sql
CREATE TABLE cloudtech_forum.posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 環境変数ファイルの更新
1. `03_api/3.2_REATful_API/02_Create/cloudtech_forum/.env`ファイルを開く

2. 下記の環境変数において、DB_PASSWORDをご自身のものに合わせて更新する

```
API_PORT=8080
DB_USERNAME=root
DB_PASSWORD=ここを更新
DB_HOST=localhost
DB_PORT=3306
DB_NAME=cloudtech_forum
```

## HTTPサーバの起動

1. 下記のコマンドでフォルダーの移動を行う
```sh
cd 03_api/3.2_REATful_API/02_Create/cloudtech_forum
```

2. Goのアプリケーションを実行し、HTTPサーバを起動する
```sh
$ go run main.go
```

3. 以下のような正常終了を示すメッセージが表示されることを確認
```sh
2025/04/15 10:21:05 APIサーバを起動しました。ポート: 8080
```

## Createの実行
1. 以下のcurlコマンドで、投稿データの登録を行う

```sh
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"content": "AWSはじめました", "user_id": 1}'
```

2. 下記のように正常終了を示すメッセージが表示されることを確認

```sh
{"id":1,"message":"登録が成功しました"}
```

## 登録データの確認

1. 以下のコマンドでMySQLにログイン
```
mysql -u root -p
```

2. 以下のコマンドで、データの検索を行う
```sql
select * from cloudtech_forum.posts;
```

3. 以下のように、登録したデータが表示されることを確認する
```
+----+-----------------------+---------+---------------------+---------------------+
| id | content               | user_id | created_at          | updated_at          |
+----+-----------------------+---------+---------------------+---------------------+
|  1 | AWSはじめました       |       1 | 2025-04-15 10:30:24 | 2025-04-15 10:30:24 |
+----+-----------------------+---------+---------------------+---------------------+
1 row in set (0.00 sec)
```