
## 1. APIサーバを作成
1. EC2のダッシュボードを開く
2. 左側のメニューから、`インスタンス`をクリック
3. 右上の`インスタンスを起動`をクリック
4. 名前：`APIServer`
5. Amazonマシンイメージ：`Amazon Linux 2023 AMI`
6. インスタンスタイプ：`t2.micro`
7. VPC：`my-vpc`
8. サブネット：`api-subnet-01`
9. パブリックIPの自動割り当て：`有効化`
10. `セキュリティグループを作成`をクリック
11. セキュリティグループ名：`api-sg`
12. `インスタンスを起動`をクリック

## 2. db-sgの設定変更（APIサーバからのMySQL通信を許可）
1. セキュリティグループの一覧から、`db-sg`を選択
2. `インバウンドルールを編集`をクリック
3. 既存のルールを`削除`
4. `ルールを追加`をクリック
5. タイプ：`MYSQL/Aurora`
6. ソース：`api-sg`
7. `ルールを保存`をクリック

## 3. APIサーバにログイン
1. 

## 4. mysqlのインストール

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

### 2. RDSに接続

以下のコマンドで、RDSのMySQLに接続
```
mysql -h testdb.co7vcyxdczfg.ap-northeast-1.rds.amazonaws.com -P 3306 -u admin -p
```

AWS RDSインスタンスのMySQLデータベースに接続します。以下のコマンドを実行前に、適切なエンドポイントとユーザ名に置き換えてください。また、パスワードの入力を求められるため、はRDSインスタンス作成時に指定したものを入力してください。


## 3. テストデータを作成
`testdb`データベースを作成する
```sql
CREATE DATABASE testdb;
```

`test_table`テーブルを作成する
```sql
CREATE TABLE testdb.test_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    age INT
);
```

テスト用のデータを登録する
```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Taro', 30);
```

```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Jiro', 22);
```

```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Hanako', 25);
```

登録されたデータの確認をする
```sql
SELECT * FROM testdb.test_table;
```

## 9. MySQLからログアウト
MySQLからログアウトする
```
exit
```


## 10. ECSタスクに環境変数を設定
1. ECSのダッシュボードを開く
2. 左側のメニューから、`タスク定義`を開く
3. `api-task`をクリック
4. 右上の`新しいリビジョンの作成`をクリック
5. 環境変数（オプション）のところにある環境変数を設定
  - DB_SERVERNAME：RDSのエンドポイント
  - DB_USERNAME：RDS作成時に指定したユーザ名（デフォルトは`admin`）
  - DB_PASSWORD：RDS作成時に指定したパスワード

## 11. api-serviceの更新
1. 左側のメニューから、`クラスター`を選択
2. `MyCluster`を選択
3. サービスから`api-service`を選択
4. 右上の`サービスを更新`をクリック
5. タスク定義のリビジョンを、（最新）とつくものに変更
6. `更新`をクリック
7. タスクが再起動されることを確認、されない場合、タスクを停止し、再起動する

## 12. 動作確認
以下のURLで、データベース接続が成功したメッセージが表示されることを確認
http://<ALBのDNSドメイン名>/dbtest