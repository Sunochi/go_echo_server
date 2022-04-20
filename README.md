# golang_echo_server
Golang echo server.

## 環境

以下の環境以外で動作確認はしていないため、動作しない可能性あり
```
MacOS: 12.3.1（Monterey, Apple M1）
Golang: 1.18.1
Docker: 20.10.14
```

## 注意

事前にDocker Desktopをインストールしたください。

`make init`を実行した時、以下のエラーが出た場合はXcodeのインストールが必要です。
```
$ make init
xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun
```

## セットアップ
```
# 環境セットアップ
$ make init
```

## サーバー起動
```
$ make up
```
※ Terminal　を占有するため、別タブで作業推奨

## サーバー再起動(Golangの修正反映)
```
$ make restart
```
もしくは `Ctrl-c + make up`

## go mod tidy実行
```
$ docker-compose exec app sh
/go/src/app # go mod tidy
```

## Lint, Format
```
$ make lint

$ make format
```

## Golang パッケージ化
```
$ make package
```

## minio
テスト環境のみExAws.S3の接続先をminioという仮想ストレージに変更しています。
S3接続をモック化せずにテストを書くことができます。
バケットとオブジェクトの内容はブラウザから確認できます。
```
$ docker-compose up -d
$ open http://localhost:9001/login
```
### ログイン情報
Username: minio_user
Password: minio_password
(※ docker-compose.yml参照)

## 動作確認URL

- CSV動作確認用URL `localhost:8080/`
  - test用のCSVは `testdata/user.csv` に格納済み

## TODO

- S3にアップロードする処理の追加
- testの追加
- `go mod tidy` のmake化