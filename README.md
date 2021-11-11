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

## Usage

```
Usage:
  pimento [command]

Available Commands:
  add         クレデンシャル情報を追加で保存します
  completion  generate the autocompletion script for the specified shell
  delete      保存されたクレデンシャル情報を削除します
  edit        対象のクレデンシャルを更新する
  help        Help about any command
  import      CSVからクレデンシャルを取り込む
  init        pimento の設定ファイルを生成する
  list        保存しているクレデンシャルの一覧を表示します.
  show        対象のクレデンシャルを1件取得する

Flags:
      --config string       config file (default is $XDG_CONFIG_HOME/pimento/config.yaml)
  -h, --help                help for pimento
      --secret_key string   pimento secret key
```
