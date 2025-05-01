# ハンズオン手順

## 1. web-subnet-01を作成
1. 左側のメニューから、`サブネット`を選択
2. 右上の`サブネットを作成`をクリック
3. VPCとして、先ほど作成した`my-vpc`を選択
4. サブネットの設定を以下のように入力
    - サブネット名：`web-subnet-01`
    - アベイラビリティゾーン：`ap-northeast-1a`
    - IPv4サブネットCIDRブロック：`10.0.3.0/24`
5.  入力が完了したら、`サブネットを作成`をクリック
6. `web-subnet-01`が無事に作成されることを確認

## 2. WEBルートテーブルを作成
1. 左側のメニューから、`ルートテーブル`を選択
2. 右上の`ルートテーブルの作成`をクリック
3. 名前に、`web-routetable`を入力
4. VPCに、my-vpcを選択
5. `ルートテーブルの作成`をクリック
6. `ルートを編集`をクリック
7. `ルートを追加`をクリック
8. 送信先に`0.0.0.0/0`、ターゲットに`my-internet-gateway`を入力し、`変更を保存`をクリック
9. 左側のメニューから、`サブネット`を選択
10. `web-subnet-01`のapi部分を選択
11. ルートテーブルのタブをクリックし、`ルートテーブルの関連付けを編集`をクリック
12. ルートテーブルIDを、`web-routetable`に変更します。
13. `保存`をクリック

## 3. EC2インスタンスを作成
1. EC2のダッシュボードを開く
2. 左側のメニューから、`インスタンス`をクリック
3. 右上の`インスタンスを起動`をクリック
4. 名前：`WebServer`
5. Amazonマシンイメージ：`Amazon Linux 2023 AMI`
6. インスタンスタイプ：`t2.micro`
7. VPC：`my-vpc`
8. サブネット：`web-subnet-01`
9. パブリックIPの自動割り当て：`有効化`
10. `セキュリティグループを作成`をクリック
11. セキュリティグループ名：`web-sg`
12. 以下を入力
    - タイプ：`HTTP`
    - ソース：`0.0.0.0/0`
13. `インスタンスを起動`をクリック

## 4. フロントアプリケーションをインストール
1. WebサーバのEC2インスタンスに、sshなどでログインする

2. システムを最新の状態に保つため、以下のコマンドでyumパッケージを更新する
    ```shell
    sudo yum update -y
    ```

3. EC2インスタンスにソースコードをダウンロードするため、Gitをインストールします。
    ```shell
    sudo yum install -y git
    ```

4. WebサーバとしてNginxをインストールし、起動します。
    ```shell
    sudo yum install -y nginx
    sudo systemctl start nginx
    sudo systemctl enable nginx
    ```

5. ブラウザで `http://[web-serverのパブリックIPアドレス]` を開き、`Welcome to nginx!` のページが表示されることを確認

6. Gitを使用してソースコードを以下のディレクトリにクローンする
    ```shell
    cd /usr/share/nginx/html/
    sudo git clone <ご自身のcloudtech-forum-webのURL>
    ```

7. Nginxの設定ファイルを編集し、表示するWebページのディレクトリを変更します。
    ```shell
    sudo vi /etc/nginx/nginx.conf
    ```
    - 変更前の設定：`root /usr/share/nginx/html;`
    - 変更後の設定：`root /usr/share/nginx/html/cloudtech-forum-web;`

8. 設定を変更した後、Nginxを再起動して変更を適用
    ```shell
    sudo systemctl restart nginx
    ```

## 5. API接続先の設定
1. WebアプリケーションからAPIを呼び出すための設定ファイル `config.js` を編集する
    ```shell
    sudo vi /usr/share/nginx/html/cloudtech-reservation-web/config.js
    ```
    baseURLの設定を、APIサーバのパブリックIPアドレスに設定値を変更します。
    ```javascript
    const apiConfig = {
    baseURL: 'http://[API-serverのパブリックIPアドレス]'
    };
    ```

## 6. 動作確認
以下のURLでWebアプリケーションが正しく起動していることを確認する
`http://[web-serverのパブリックIPアドレス]`