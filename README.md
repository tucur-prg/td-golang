td-golang
----

https://go.dev/

https://go.dev/doc/effective_go


コンテナ実行環境
```
docker run --rm -it -v `pwd`:/app -w /app -p 9090:9090 golang:1.23-bullseye bash
```

https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode

```
# パッケージのキャッシュ
$GOPATH/pkg

# ビルドのキャッシュ
go env | grep GOCACHE
~/.cache/go-build
```
