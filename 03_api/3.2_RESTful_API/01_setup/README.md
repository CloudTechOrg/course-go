# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- Gitがインストールされていること
- GitHubアカウントが作成されていること

# ハンズオン手順

## 1. MySQLのデータベース作成
1. 以下のコマンドでMySQLにログイン
    ```
    mysql -u root -p
    ```

2. 以下のコマンドで、`cloudtech_forum`のデータベースを作成する
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
2. 右上の`+`ボタンから、`New repository`を選択する
3. Repository Nameに`cloudtech-forum`を入力する
4. Public or Privateは任意（迷う場合はPublicでOK）
5. 入力が完了したら`Create Repository`をクリック    
6. 無事にリポジトリが作成されれば、ここまでの操作は完了

## 3. Visual Studio Codeのフォルダー作成
1. ご自身のPC内の任意の箇所に、`cloudtech-forum`というフォルダーを作る
2. Visual Studio Codeを開き、さきほど作成した`cloudtech-foru`のフォルダーを開く
3. `README.md`というファイルを作成（該当リポジトリの説明や使い方を記載するファイル、中身はいったん空でOK）
4. `.env`というファイルを作り、下記内容を書き込む（環境変数を設定するファイル、パスワードはご自身のMySQLのパスワードを設定）
    ```
    API_PORT=8080
    DB_USERNAME=root
    DB_PASSWORD=<ご自身のパスワードを設定>
    DB_HOST=localhost
    DB_PORT=3306
    DB_NAME=cloudtech_forum
    ```
4. `.gitignore`というファイルを作り、下記内容を記載（機密情報を含むファイルなどがGitHubにアップロードされるのを防ぐファイル）
    ```.gitignore
    .env
    *.log
    *.tmp
    *.db
    vendor/
    ```

## 4. モジュールの作成
1. 以下のコマンドで、`cloudtech-forum`という名称のGoのモジュールを作る
    ```
    go mod init cloudtech-forum
    ```
2. `go.mod`ファイルが作成されることを確認


## 5. Gitの初期設定
1. Visual Studio Codeのターミナルを開く
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
6. リモートブランチとして、さきほど作成したcloudtech_forumのリポジトリを指定（<your-github-repository-url>はご自身のものに置き換え）
    ```shell
    git remote add origin <your-github-repository-url>
    ```
7. 変更内容をGitHubに反映
    ```shell
    git push origin main
    ```
8. GitHubの該当リポジトリに、`.gitignore`と`README.md`がアップロードされ、`.env`はアップロードされていないことを確認