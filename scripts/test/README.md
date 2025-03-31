
https://pkg.go.dev/testing

go test は実行するとテスト結果をキャッシュして、変更がなければその部分のテストを再実行しない  
キャッシュをさせなくする場合は、「-count=1」をオプションとしてつける  
オプションの用途はキャッシュさせなくするためのものではなく、テストの実行回数を指定するもの  

```sh
# 全体テスト
go test ./...

# 関数指定
go test ./... -run <MethodName>
```

[テスト（go test/testing）｜Go言語入門](https://www.twihike.dev/docs/golang-primer/testing)  

testing パッケージを使ったテストのやり方

テスト用のソースファイルを作成します。
* ファイル名の末尾は_test.goとする
* パッケージ名はテスト対象と同じにする（別パッケージも可、以降で説明）
* テスト関数のシグネチャはfunc TestXxx(t *testing.T)とする（Xxxは小文字で始まらない）
* testing.T.Errorfを使ってログ出力とテストの失敗報告をする

[Go言語でユニットテスト, テストしやすいコードとモックを書く #mockgen - Qiita](https://qiita.com/hiroyky/items/4a9be463e752d5c0c41c)  

[テストデータをどこに配置するべきか | Go言語 - BioErrorLog Tech Blog](https://www.bioerrorlog.work/entry/golang-test-data-location)  

> testdata という名のディレクトリを無視するのでここにおく

# benchmark


# fuzzing

[GoのFuzzingテストやってみる](https://zenn.dev/10inoino/articles/fc2551b3b3355e)  

> １つ注意点として、Fuzzingテストはテストが失敗するまで実行し続けるので、タイムアウトの時間を設定して実行するのが良いでしょう。


# ビジュアライズ

https://github.com/mfridman/tparse

[Goのテスト結果をtparseで整形する・GitHub ActionsのJob Summaryと組み合わせる - 私が歌川です](https://blog.utgw.net/entry/2023/06/16/162330)  

```sh
go test ./... -json | tparse -all
```

# error

## $WORK/b001/_testmain.go:14:8: could not import main (cannot import "main")

[go mod init mainをするとテストコードは動かない](https://zenn.dev/axpensive/articles/c9b85f7559a235)  

> go mod initでmodule名をmainにしたのが原因

