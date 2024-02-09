# Golangの開発環境について

## goプロジェクトを作成する

1. `go mod init {モジュールパス}` コマンドを実行し、`go.mod`ファイルを作成する。モジュールパスに指定したものは、go.modのmoduleに指定される

    例: `go mod init github.com/yusuke-rakus/example_app`
    ```
    module github.com/yusuke-rakus/example_app

    go 1.21.6
    ```

1. `main.go`などのプロジェクトを作成する
    ```go: main.go
    package main

    import (
        "net/http"
        
        "github.com/labstack/echo/v4"
    )

    func main() {
        e := echo.New()
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Hello, World!")
        })
        e.Logger.Fatal(e.Start(":1323"))
    }    
    ```

1. `go mod tidy`を実行し、

    - 不足しているパッケージがあれば`go.mod`に追加
    - `go.mod`に記載されているパッケージをインストールし、`go.sum`を作成する


## インストールコマンド
`go get *`と`go install *`があるが、最近はinstallの方が推奨されているらしい

ただし、プロジェクトで使うものは`go.mod`で管理したいため、パッケージのインストール系は後述する`go mod tidy`で行いたい


## `go mod tidy`の使い方

1. 開発中、importするパッケージを記述した後、コマンド実行

    →`go.mod`にパッケージを追加、`go.sum`にも追加

1. プロジェクトからimportを削除後、コマンド実行

    →`go.mod`からパッケージを削除、`go.sum`からも削除

> tidyとは、「こぎれいな」みたいな意味らしい
> 
> つまりソースコードの状況を見ていい感じにgo.sumとgo.modを編集してくれるみたい