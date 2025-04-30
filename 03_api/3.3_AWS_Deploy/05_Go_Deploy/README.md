# 前提事項
- AWSアカウントが作成されており、AWSにログインしていること
- APIサーバのEC2インスタンスにssh接続などでログインしていること

# 手順
# 1. データベースの設定
1. パッケージマネージャー（yum）の更新を行う
```shell
sudo yum update -y
```

2. MySQLの公式リポジトリをシステムに追加する
```shell
sudo yum install https://dev.mysql.com/get/mysql84-community-release-el9-1.noarch.rpm
```

3. MySQLサーバのインストールを行う
```shell
sudo yum install mysql-community-server -y
```

4. MySQLサービスの起動を行う
```shell
sudo systemctl start mysqld
```

5. システム起動時にMySQLが自動起動するように設定する
```shell
sudo systemctl enable mysqld
```

6. 以下のコマンドで、RDSのMySQLに接続する
```
mysql -h <RDSエンドポイント> -P 3306 -u admin -p
```

AWS RDSインスタンスのMySQLデータベースに接続します。以下のコマンドを実行前に、適切なエンドポイントとユーザ名に置き換えてください。また、パスワードの入力を求められるため、はRDSインスタンス作成時に指定したものを入力してください。

7. 以下のコマンドで、`cloudtech_forum`のデータベースを作成
```sql
CREATE DATABASE cloudtech_forum;
```

8. 以下のコマンドで、postsのテーブルを作成する
```sql
CREATE TABLE cloudtech_forum.posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

9. MySQLからログアウト
MySQLからログアウトする
```
exit;
```

## 2. Goアプリケーションの設定
1. システムを最新の状態に保つためにyumパッケージをアップデートする
```shell
sudo yum update -y
```

2. EC2インスタンスにソースコードをダウンロードするために、Gitをインストールする
```shell
sudo yum install -y git
```

3. APIサーバとして機能するGo言語をインストールする
```shell
sudo yum install -y golang
```

4. インストール後、Goのバージョンを確認する
```shell
go version
```

5. Gitを使用してソースコードをダウンロードする
```shell
cd /home/ec2-user/
git clone https://github.com/CloudTechOrg/cloudtech-forum.git
```

## 3. 環境変数の設定
GoのアプリケーションからRDSに接続するための設定ファイルを作成します。

1. `vi`コマンドで`.env`ファイルを作成する

```shell
vi cloudtech-forum/.env
```

2. 以下の内容を記載する
```
API_PORT=8080
DB_USERNAME=admin
DB_PASSWORD=【RDSインスタンスのパスワード】
DB_HOST=【RDSインスタンスのエンドポイント】
DB_PORT=3306
DB_NAME=cloudtech_forum
```

## 4. サービスの自動起動設定
システムの再起動時にもAPIが自動で起動するようにsystemdを設定します。

1. viエディターを使用し、サービス起動時の設定ファイルを作成する
```shell
sudo vi /etc/systemd/system/goserver.service
```

2. 以下の内容をファイルに追記し、保存を行う
```
[Unit]
Description=Go Server

[Service]
WorkingDirectory=/home/ec2-user/cloudtech-forum
ExecStart=/usr/bin/go run main.go
User=ec2-user
Restart=always

[Install]
WantedBy=multi-user.target
```

設定を有効にし、サービスを開始します。
```shell
sudo systemctl daemon-reload
sudo systemctl enable goserver.service
sudo systemctl start goserver.service
```

## 5. 動作確認
1. 以下のコマンドで、postデータの登録を行う

```sh
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"content": "AWSはじめました", "user_id": 1}'
```

2. 以下のコマンドで、postデータの登録確認を行う
```sh
curl -X GET http://localhost:8080/posts
```

## 6. リバースプロキシの設定
8080ポートで動作するGoのAPIを80ポートで利用できるように、Nginxをリバースプロキシとして設定します。
```shell
sudo yum install nginx
sudo systemctl start nginx
sudo systemctl enable nginx
```
Nginxの設定ファイルを編集し、適切なリバースプロキシ設定を行います。
```shell
sudo vi /etc/nginx/nginx.conf
```

`server { ・・・ }` の部分を、下記内容に変更します
```
server {
        listen 80;
        server_name _;
        location / {
            proxy_pass http://localhost:8080;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host $host;
            proxy_cache_bypass $http_upgrade;
        }
    }
```


設定を更新した後、Nginxを再起動します。
```shell
sudo systemctl restart nginx
```

# 動作確認

## 80ポートでの起動確認
```sh
curl -X GET http://localhost/posts
```

## 外部からのアクセス
以下のアドレスをブラウザにて実行する
- `http://[EC2インスタンスのパブリックIPアドレス]/posts`