# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること

# ハンズオン手順

## 1. MySQLのデータベース作成
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

## 2. GitHubリポジトリの作成
1. GitHubアカウントにログインする
2. 右上の`+`ボタンから、`New repository`を選択します。
3. Repository Nameに`cloudtech-forum`を入力
4. Public or Privateは任意（迷う場合はPublicでOK）
5. 入力が完了したら`Create Repository`をクリック    
6. 無事にリポジトリが作成されれば、ここまでの操作は完了

## 3. VS Codeのフォルダー作成
1. ご自身のPC内の任意の箇所に、`cloudtech-forum`というフォルダーを作る
2. VS Codeを開き、さきほど作成した`cloudtech-foru`のフォルダーを開く
3. `README.md`というファイルを作成（該当リポジトリの説明や使い方を記載するファイル、中身はいったん空でOK）
4. `.env`というファイルを作り、下記内容を書き込む（環境変数を設定するファイル、パスワードはご自身のパスワードを設定）
    ```
    API_PORT=8080
    DB_USERNAME=root
    DB_PASSWORD=<ご自身のパスワードを設定>
    DB_HOST=localhost
    DB_PORT=3306
    DB_NAME=cloudtech_forum
    ```
4. `.gitignore`というファイルを作り、下記内容を記載（機密情報を含むファイルなどがGitHubにアップロードされるのを防ぎます）
    ```.gitignore
    .env
    *.log
    *.tmp
    *.db
    vendor/
    ```

## 4. Gitの初期設定
1. VS Codeのターミナルを開く
2. 下記コマンドで、Gitの初期化を行う
    ```shell
    git init
    ```
3. ファイルの変更をステージに反映する
    ```shell
    git add .
    ```
4. 変更をコミットする
    ```shell
    git commit -m "initial commit"
    ```
5. デフォルトブランチの名前を`main`に変更
    ```shell
    git branch -m main
    ```
6. リモートブランチとして、さきほど作成したcloudtech_forumのリポジトリを指定
    ```shell
    git remote add origin <your-github-repository-url>
    ```
7. 変更内容をGitHubに反映
    ```shell
    git push origin main
    ```
8. GitHubの該当リポジトリに、`.gitignore`と`README.md`がアップロードされ、.envはアップロードされていないことを確認