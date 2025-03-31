
# constructor

構造体の場合はnew関数を使うよりも、構造体リテラルを使用してインスタンスを直接初期化する方が一般的です。

new関数を使用するとヒープメモリにインスタンスが割り当てられます。

大量の小さなオブジェクトを頻繁に生成する場合は、メモリ使用量やガベージコレクションに影響を与える可能性があるので注意しながら実装する必要があります。

[Go言語で完全コンストラクタパターンを実装する #PHP - Qiita](https://qiita.com/kbys-fumi/items/669db4385c3e8471088a)

https://go.dev/doc/effective_go#composite_literals

https://stackoverflow.com/questions/34774615/should-a-constructor-function-return-an-error-or-a-null-value


# receiver

[Go言語のポインタレシーバと値レシーバの使い分け #Go - Qiita](https://qiita.com/atsutama/items/32a3961e1e74e20bcb14)    
[【Go言語入門】Goのレシーバーとは？2種類のレシーバーを使い分ける方法 #Go - Qiita](https://qiita.com/ryo_manba/items/c567858befd04602e3ec)


# embedded field

[Go言語で「embedded で継承ができる」と思わないほうがいいのはなぜか？ #オブジェクト指向 - Qiita](https://qiita.com/Maki-Daisuke/items/511b8989e528f7c70f80)  


# tag

[Go言語(golang) reflectで構造体のタグ/アノテーションを取得する - golangの日記](https://golang.hateblo.jp/entry/2018/11/10/084500)  
[Go リフレクション入門 #Go - Qiita](https://qiita.com/s9i/items/b835634d84bba5574d0a#%E3%82%BF%E3%82%B0%E6%83%85%E5%A0%B1%E3%81%AE%E5%8F%96%E5%BE%97)  


# private field

[Goで構造体の非公開フィールドにアクセスする方法 - stop-the-world](https://stop-the-world.hatenablog.com/entry/2019/12/31/214058)  

> Go の 構造体 (struct) におけるフィールドは、フィールド名が小文字始まりであれば 非公開フィールド (unexported field) となり、パッケージ外からアクセスすることができません
>
> https://go.dev/ref/spec#Exported_identifiers



