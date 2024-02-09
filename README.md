# GoでEchoフレームワークを用いて開発するメモ

## 開発環境に関するメモ
- 使用エディタ: VS Code
- Docker上で実行し、そこにリモート接続し開発を行う(DevContainer)
- DB: MySql(別コンテナを作成する)
- Airでホットリロードが使えるため、コードを変更しても再ビルドを自動でやってくれて便利
- 何かフォーマットしてくれるようなライブラリなど欲しい

## Echoに関するメモ

## 実行コマンド
```
$ air -c .air.toml // airのパスが設定されてないとエラーになる
$ /go/bin/linux_amd64/air -c .air.toml
```


## go.modの更新
```
$ go mod tidy
```

## Airが効いている風なのに変更が反映されない
`command + Shift + P`でコマンドパレットを開き、`rebuild`と入力し`開発コンテナー：キャッシュなしのコンテナーのリビルド`を実行