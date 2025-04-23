# 概要
この説明は、ハンズオン課題における`Sprint1：AWS基本サービス1`にて必要となる、APIサーバに対する設定方法を説明しています。

# 前提
- APIサーバのEC2インスタンスにssh接続などでログインしていること

# 手順

## 1. yumのアップデート
システムを最新の状態に保つためにyumパッケージをアップデートします。
```shell
sudo yum update -y
```

## 2. Gitのインストール
EC2インスタンスにソースコードをダウンロードするために、Gitをインストールします。
```shell
sudo yum install -y git
```

## 3. Goのインストール
APIサーバとして機能するGo言語をインストールします。
```shell
sudo yum install -y golang
```
インストール後、Goのバージョンを確認します。
```shell
go version
```

## 4. ソースコードのダウンロード
Gitを使用してソースコードをダウンロードします。
```shell
cd /home/ec2-user/
git clone https://github.com/CloudTechOrg/cloudtech-reservation-api.git
```

## 5. サービスの自動起動設定
システムの再起動時にもAPIが自動で起動するようにsystemdを設定します。

まずはviエディターを使用し、サービス起動時の設定ファイルを作成します。
```shell
sudo vi /etc/systemd/system/goserver.service
```
以下の内容をファイルに追記し、保存を行います。
```
[Unit]
Description=Go Server

[Service]
WorkingDirectory=/home/ec2-user/cloudtech-reservation-api
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
## 8080ポートでの起動確認
```
curl http://localhost:8080
```

## 80ポートでの起動確認
```
curl http://localhost
```

## 外部からのアクセス
以下のアドレスをブラウザにて実行する
- `http://[EC2インスタンスのパブリックIPアドレス]`