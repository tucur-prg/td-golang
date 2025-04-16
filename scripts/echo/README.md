
[High performance, extensible, minimalist Go web framework | Echo](https://echo.labstack.com/)

[echo package - github.com/labstack/echo/v4 - Go Packages](https://pkg.go.dev/github.com/labstack/echo/v4)

# https

[Go による TLS 通信(サーバ側)](https://zenn.dev/empenguin/articles/5a8a8d827cfb1c)


# echo.Context

[Binding | Echo](https://echo.labstack.com/docs/binding)  
[Echo(Go)で可変のJSONをBINDする](https://zenn.dev/kaikusakari/articles/becace7d43abb5)  

[Request | Echo](https://echo.labstack.com/docs/request#validate-data)  
[【Golang】echoとvalidator.v9を使ったカスタムバリデーションのサンプル #Go - Qiita](https://qiita.com/nanamen/items/c824a2c8f2e1767f90f8)  
[【Go】validatorのタグでカスタムメッセージを設定する #Go - Qiita](https://qiita.com/yudai2929/items/a0b0213b3f8b0a459f44)  


# echo.Route

[Routing | Echo](https://echo.labstack.com/docs/routing)

> 各登録メソッドはRoute、登録後にルートに名前を付けるために使用できるオブジェクトを返します。
> ルート名は、テンプレートから URI を生成するとき、ハンドラー参照にアクセスできないとき、または同じハンドラーを持つ複数のルートがあるときに非常に役立ちます。


echoでルーティングからURLを生成する場合は「c.Echo().Reverse(Name, Params ...)」を使う

```go
e.GET("/users/:id", h).Name = "foobar"

func h(c echo.Context) error {
  url := c.Echo().Reverse("foobar", "123")
  return c.String(http.StatusOK, url)
}
```

> Globalにキャッシュされてそうな値をリクエスト単位にする方法はあるのか？
> 
> テナント型の構造でAPIリクエストのクライアントクレデンシャルがタイミングによって別のテナントのものになるのでは？
> 
> 
> ルーティング設定する時にクラスで渡した場合に、クラス内の変数が状態を保持しつづけてしまうので前回の状態を残したまま次のリクエストを受け付けてしまう
> 
> 状態を受け継ぎたくないものはローカル変数にしないといけない


# Middleware

[Middleware | Echo](https://echo.labstack.com/docs/cookbook/middleware)  

[[Golang]EchoのMiddlewareを実装する - ken-aio's blog](https://ken-aio.github.io/post/2019/02/06/golang-echo-middleware/)  

[Goを学ぶ　～echoのミドルウェア機能を見る～ part3 #Docker - Qiita](https://qiita.com/dsricekun/items/81d6bf78ea5b3ba7c418)  

[Try Golang! EchoでオリジナルのMiddlewareを作ろう！ | by Takuo | VELTRA Engineering | Medium](https://medium.com/veltra-engineering/echo-middleware-in-golang-90e1d301eb27)  
> Pre, Use, Group, Routeの4通りの方法でMiddlewareを設定しています。


ミドルウェアからデータを渡す時はc.Set(key, val)を使って渡す

# ACL

[Casbin Auth | Echo](https://echo.labstack.com/docs/middleware/casbin-auth)  

[Casbinで始めるアクセス制御 | フューチャー技術ブログ](https://future-architect.github.io/articles/20221004a/)  

# Error Handling

[Error Handling | Echo](https://echo.labstack.com/docs/error-handling)

# profiling

## pprof

https://pkg.go.dev/net/http/pprof

https://github.com/google/pprof#basic-usage

[Go の net/http/pprof で CPU や Memory を profile する流れと内部実装 - sambaiz-net](https://www.sambaiz.net/article/238/)

```sh
docker run --rm -it golang:1.23-bullseye bash

apt-get update && apt install -y graphviz

go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
```

[golang の net/http/pprof を触ってみたメモ - sonots:blog](https://sonots.livedoor.blog/archives/39879160.html)  
> blocking profiler を有効にするために runtime.SetBlockProfileRateを呼んでいる。

[pprof について](https://zenn.dev/muroon/articles/adf577f563c806)

## trace

[golang でのデバッグに非常に便利な go tool trace - 理系学生日記](https://kiririmode.hatenablog.jp/entry/20190506/1557097529)  


# etc

[Static | Echo](https://echo.labstack.com/docs/middleware/static)  
> 静的ファイルの扱い

[Graceful Shutdown | Echo](https://echo.labstack.com/docs/cookbook/graceful-shutdown)  

[echoでgraceful restart - DevDevデブ!!](https://stk132.hatenablog.com/entry/2017/09/26/022021)  


