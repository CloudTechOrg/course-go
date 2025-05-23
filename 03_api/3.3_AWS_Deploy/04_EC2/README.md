
# 前提事項
- AWSアカウントが作成されており、AWSにログインしていること

# ハンズオン手順
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
12. 以下を入力
    - タイプ：`HTTP`
    - ソース：`0.0.0.0/0`
13. `インスタンスを起動`をクリック

## 2. RDSに接続できるように設定
1. 左側のメニューから、`セキュリティグループ`を選択
2. `db-sg`を選択
3. `インバウンドルールを編集`をクリック
4. `ルールを追加`をクリック
5. 以下内容を設定
    - タイプ：`MYSQL/Aurora`
    - ソース：`api-sg`
6. `ルールを保存`をクリック