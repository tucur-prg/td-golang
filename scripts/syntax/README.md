
[とほほのGo言語入門 - とほほのWWW入門](https://www.tohoho-web.com/ex/golang.html)  

| 演算子 | 説明 |
|----|----|
| x := y | xにyを代入(初期化の使用可能) 型推論 |
| ch <- x | chチャネルにxを送信 |
| x =<- ch | chチャネルからxに受信 |

[Google Go Style Guideを読んでみよう！ - Style Decisions編 #Styleguide - Qiita](https://qiita.com/TakumaKurosawa/items/fbb1418111604837d8ac)  


[【Go】シングルクォーテーションとダブルクォーテーションの型 #初心者 - Qiita](https://qiita.com/obr-note/items/a3b81e258494dbf470b3)  

> シングルクォーテーションはrune型、ダブルクォーテーションはstring型で違う型になる

[ゼロ値を使おう #golang #Go - Qiita](https://qiita.com/tenntenn/items/c55095585af64ca28ab5)  

> Goの変数は必ず初期化される

[Golang の変数キャプチャ | ツイートするには長すぎる](https://blog.nfurudono.com/posts/go-learn-memory-motivation/)  


# init

[[golang]初期化関数initについて #Go - Qiita])(https://qiita.com/tom-takeru/items/824ce5c12a59d7ba5963)  

# block

[Go言語におけるブロックとスコープ](https://zenn.dev/dqneo/articles/9e8a0d6f67f8f9510891)  

# pointer

型のプレフィックスで * をつけるとポインタ型を定義できる  
値を入れるときにプレフィックスで & をつけるとポインタ型を生成できる  

[Goで学ぶポインタとアドレス #プログラミング - Qiita](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)  
[Goにおけるポインタの使いどころ](https://zenn.dev/uji/articles/f6ab9a06320294146733)  
[【Go】ポインタっていつ使うん？ #AdventCalendar2023 - Qiita](https://qiita.com/nakampany/items/309174d299a69738179b)  
[[Go] Pointer](https://zenn.dev/yagi_eng/scraps/48771c75c08afb)  

[リテラル値のポインタ](https://zenn.dev/spiegel/articles/20211004-pointer-to-literal-value)  

> リテラル表現から直接ポインタ値を得ることもできない。

[Goでよく見るnewとNewを調べたときのメモ #Go - Qiita](https://qiita.com/gold-kou/items/4494f8b69b8fa53d5e93)  

> newは指定した型のポインタ型を生成する関数です。

[【Go】 make と new の違い #Go - Qiita](https://qiita.com/kei3dev/items/da7f8d54036753c6e473)  

# Variable

## const

[Go言語の定数を解説 - Recursion](https://recursionist.io/learn/languages/go/data-type/constant)  

[[Golang] 定数 - Golangで定数を宣言して使う方法について説明します。](https://deku.posstree.com/golang/constants/)  

## global

[Goのグローバル変数とスコープでハマった話 #テスト - Qiita](https://qiita.com/UHNaKZ/items/637cb3e1c538d8e63ee2)  

> 初期化代入（:=）するとグローバルじゃなくてローカル変数になってしまう


## Named return values

https://go.dev/doc/effective_go#named-results

https://go.dev/tour/basics/7

[GoのNamed return valueについてメリデメを考える](https://zenn.dev/yuyu_hf/articles/c7ab8e435509d2)  


# defer

[Goのdefer文を使うときに気をつけること - Yappli Tech Blog](https://tech.yappli.io/entry/understanding-defer-in-go)  

> defer文はifやforなどの制御構文の一種で、以下のように使うことで与えられた処理を関数がreturnされた後や関数の末尾に到達した後に実行させることができます。


# channels

[Go 言語のチャネル - Go言語の基礎 - Go 言語入門](https://golang.keicode.com/basics/go-channel.php)  

[Go で１行でgoroutineデッドロックを起こす方法](https://zenn.dev/zawawahoge/articles/fd533a54ee6e2b)  


# switch

[Go言語(golang) switch文の書き方 - golangの日記](https://golang.hateblo.jp/entry/2019/10/07/225026)  

# cast / assertion

[Goの型変換とアサーションの違い](https://zenn.dev/the_exile/articles/494090c0822a1a)  

> 型変換とは、ある型の値を別の型に変換することです。

> アサーションとは、インターフェース型の値が実際にはどの具体的な型を持っているかを確認したり、その具体的な型として扱ったりすることです。

## 文字列変換

https://pkg.go.dev/strconv

[Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](https://matope.hatenablog.com/entry/2014/04/22/101127)  

# panic / recover

[Go言語のパニックとリカバーを解説 - Recursion](https://recursionist.io/learn/languages/go/error/panic-recover)  



# 組み込み

## embed

[Go 1.16からリリースされたgo:embedとは | フューチャー技術ブログ](https://future-architect.github.io/articles/20210208/)  

> go:embedでは一見コメントアウトに見える//go:embed sample.jsonが埋め込みファイルの場所を指示する記述として機能します。

[go:embedを使ってローカルにあるファイルを呼び出す方法](https://zenn.dev/rescuenow/articles/aeb7f2e8c110d0)  

## build

ビルド時にタグが入っている条件に一致しない場合は読み込まれない

[Goのビルドタグの書き方が// +buildから//go:buildに変わった理由](https://zenn.dev/team_soda/articles/golang-build-tags-history)  

[【Go言語】ビルドタグgo:buildの使い方 #Go - Qiita](https://qiita.com/twrcd1227/items/f5d787e22d2379baca28)  

## Context

[Go の Context を学ぶ #Go - Qiita](https://qiita.com/TsuyoshiUshio@github/items/34b63b663ffd56125c07)  

> ざっとまとめると、Context は、APIのサーバーやクライアントを使うときに、コンテキストを提供してキャンセルや、タイムアウト、値を渡したり出来る仕組み。

[golang contextの使い方とか概念(contextとは)的な話 #Go - Qiita](https://qiita.com/marnie_ms4/items/985d67c4c1b29e11fffc)  

> **構造体の中には、Contextを保持してはならない**  
> 各Contextに基づく処理(キャンセル・値取得)はリクエストスコープのみに影響範囲がとどまるべきである。という点からです。  
> 例えば、各リクエストで共有されるようなグローバルスコープの構造体に保持する事によって  
> 意図せず別のリクエストによる副作用を受ける可能性があるので、そもそも保持しない事でそれを回避したほうが良い、という事です。  

[初心者がGo言語のcontextを爆速で理解する ~ value編　~ #新人プログラマ応援 - Qiita](https://qiita.com/yoshinori_hisakawa/items/50966e9ba2627e5ac124)  

> func WithValue(parent Context, key interface{}, val interface{}) Context
> 
> 第一引数に、key-valueを格納するcontextを指定し、key-valueをセットします。
そしてセット済みのcontextが返却されます。
> 
> 親子関係になる？親の値を受け継いでいく

[Valueメソッドを有効に使うtips｜よくわかるcontextの使い方](https://zenn.dev/hsaki/books/golang-context/viewer/appliedvalue)  

