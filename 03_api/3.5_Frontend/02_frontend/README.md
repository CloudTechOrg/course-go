# ハンズオン手順

## 1. GitHubリポジトリのフォーク
フロントエンドのリポジトリ（`cloudtech-forum-front`）を、ご自身のGitHubアカウントにFork（コピー）します。

1. 以下のリポジトリを開く
    https://github.com/CloudTechOrg/cloudtech-forum-front

2. 右上の`fork` → `Create a new fork`をクリック

3. Ownerからご自身のアカウントを選択し、`Create fork`をクリック

4. ご自身のアカウントに、`cloudtech-forum-front`のリポジトリがコピーされていることを確認

## 2. baseURLの変更
フロントエンドのアプリケーションが参照する、APIサーバのIPアドレスを変更します。

1. 以下のコマンドで、ローカルPCの任意の場所にダウンロードする
    ```shell
    git clone <ご自身のリポジトリのURL>
    ```

2. `config.js`ファイルを開く

3. baseURLの部分の<your-domain>の部分を、APIサーバのパブリックIPアドレスに変更する

4. 以下のコマンドで、Gitの変更を反映させる
    ```shell
    git add .
    git commit -m "update baseURL"
    git push origin main
    ```

## 3. 動作確認

1. index.htmlファイルの絶対パスをコピーし、Google Chromeにて実行する

2. トップページが表示されるが、一覧の部分になにも表示されていないことを確認する

3. Google Chromeで開発者ツールを開くと、CORSエラーが発生していることを確認する

現時点で、API側にCORS対応がされていないため、エラーが発生します。次の講座で、CORS対応を行っていきます。