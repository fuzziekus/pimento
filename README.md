# pimento
TUI PWマネージャ


## Description
TUI で利用するためのパスワードマネージャです。  
sqlite3 にクレデンシャルを保存し、sqliteのDBファイル自身をDB Keyで暗号化します。  
`pimento add` コマンドを利用することで、管理対象のクレデンシャル情報を追加します。  
このとき入力するPWは, SECRET_KEY で暗号化されてDBに格納されます。  

SECRET_KEY の優先順位は以下の通りです。
1. --secret_key オプション
2. PIMENTO_SECRET_KEY 環境変数
3. コンフィグファイルの `secret_key` 変数

