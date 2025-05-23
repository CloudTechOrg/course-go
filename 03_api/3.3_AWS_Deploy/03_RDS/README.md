# 前提事項
- AWSアカウントが作成されており、AWSにログインしていること

# ハンズオン手順
## 1. サブネットグループを作成
1. RDSのダッシュボードを開く
2. 左側のメニューから`サブネットグループ`を選択
3. 右上の`DBサブネットグループを作成`をクリック
4. サブネットグループの詳細は以下の通り入力
    1. 名前：`my-subnet-group`
    2. 説明：`my subnet group`
    3. VPC：`my-vpc`を選択
    4. アベイラビリティゾーン：`ap-northeast-1a`、`ap-northeast-1c`を選択
    5. サブネット：さきほど作った`db-subnet-01`と`db-subnet-02`を選択
6. `作成`をクリック
7. サブネットグループが正しく作成されることを確認
    
## 2. RDSインスタンスを作成
1. 左側のメニューから、`データベース`を選択する
2. 右上の`データベースの作成`をクリックする
3. データベースの作成方法は、`標準作成`を選択
4. エンジンのタイプは、`MySQL`を選択する
5. テンプレートは、`無料利用枠`を選択する
6. 可用性と耐久性はとくに変更不要（無料利用枠のため変更できない）
7. 設定は以下の通りに選択する
    1. DBインスタンス識別子：`testdb`
    2. マスターユーザー名：`admin`
    3. 認証情報管理：`セルフマネージド`
    4. マスターパスワード：任意
8. インスタンスの設定は、とくに変更なし
9. ストレージはとくに変更不要
10. 接続は以下のように設定
    1. コンピューティングリソース：`EC2コンピューティングリソースに接続しない`
    2. VPC：`my-vpc`
    3. DBサブネットグループ：`my-subnet-group`
    4. パブリックアクセス：`なし`
    5. VPCセキュリティグループ：`新規作成`
    6. 新しいVPCセキュリティグループ名：`db-sg`
    7. アベイラビリティーゾーン：`ap-northeast-1a`
    8. RDS Proxy：`チェックしない`
    9. 認証機関：`デフォルト`
11. タグはとくに設定不要
12. データベース認証は`パスワード認証`を選択
13. モニタリング以下は設定不要
14. `データベースの作成`をクリック
15. データベースが正常に作成されることを確認
