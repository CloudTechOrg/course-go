# 前提事項
- Goがインストールされていること
- MySQLがインストールされていること
- GitHubにある[https://github.com/CloudTechOrg/course-go](https://github.com/CloudTechOrg/course-go)の資材がローカルにダウンロードされていること

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
2. 右上の+ボタンから、`New repository`を選択します。
3. Repository Nameに`cloudtech_forum`を入力
4. Public or Privateは任意（迷う場合はPublicでOK）
5. 入力が完了したら`Create Repository`をクリック    
6. 無事にリポジトリが作成されれば、ここまでの操作は完了

## 3. VS Codeのフォルダー作成
1. ご自身のPC内の任意の箇所に、`cloudtech_forum`というフォルダーを作る
2. VS Codeを開く
3. VS Codeにて、さきほど作成した`cloudtech_forum`のフォルダーを開く
4. `.gitignore`というファイルを作り、下記内容を記載
    ```.gitignore
    .env
    *.log
    *.tmp
    *.db
    vendor/
    ```
5. README.mdというファイルを作成（中身はいったん空でOK）
6. VS Codeのターミナルを開く
7. 下記コマンドで、Gitの初期化を行う
    ```shell
    git init
    ```
8. ファイルの変更をステージに反映する
    ```shell
    git add .
    ```
9. 変更をコミットする
    ```shell
    git commit -m "initial commit"
    ```
10. デフォルトブランチの名前を`main`に変更
    ```shell
    git branch -m main
    ```
11. リモートブランチとして、さきほどさくせいしたcloudtech_forumのリポジトリを指定
    ```shell
    git remote add origin <your-github-repository-url>
    ```
12. 変更内容をGitHubに反映
    ```shell
    git push origin main
    ```