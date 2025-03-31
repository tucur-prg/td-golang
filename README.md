td-golang
----

https://go.dev/

https://go.dev/doc/effective_go

---

コンテナ実行環境
```
docker run --rm -it -v `pwd`:/app -w /app -p 9090:9090 golang:1.23-bullseye bash
```

[【備忘】Go 言語のビルドと実行（３パターン） #初心者 - Qiita](https://qiita.com/t-yama-3/items/1b6e7e816aa07884378e)  

```sh
# 実行
go run main.go

# 実行形式にビルド
go build main.go
./main
```

[【環境変数】go getでパッケージがインストールされるディレクトリは$GOPATH/bin。設定していない場合は$HOME/go/binに入る。 #Go - Qiita](https://qiita.com/machbike/items/f0f8c3a65c2cc6b6009e)  

[Go のモジュール管理【バージョン 1.17 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)  

```sh
# パッケージのキャッシュ
$GOPATH/pkg

# ビルドのキャッシュ
go env | grep GOCACHE
~/.cache/go-build
```
