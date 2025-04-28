

[gorm package - gorm.io/gorm - Go Packages](https://pkg.go.dev/gorm.io/gorm)  

# gorm

## Debug

import に `gorm.io/gorm/logger` を追加  
DBをオープンするときに logger のレベルを変える  

```go
	db, err := gorm.Open(sqlite.Open("development.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
```

## Migration

[Migration | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/ja_JP/docs/migration.html)  
[Golang – GORMでテーブル生成する – 諦めなければできる – Anveloper](https://anveloper.com/2020/10/08/golang-gorm/)  
[Gormのオートマイグレーション注意点 #GORM - Qiita](https://qiita.com/sky0621/items/68e59d144cfad91f27fd)  

## Preloading

[Goでよく使われるgormを理解する：Preloading編 #SQL - Qiita](https://qiita.com/tsubasaozawa/items/ac5a8a515fe4f7a139b0)  

## Model

[モデルを宣言する | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/ja_JP/docs/models.html)  
[GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】 #Go - Qiita](https://qiita.com/gold-kou/items/45a95d61d253184b0f33)  

## Error Handling

errors.Is で判別できる一覧

[gorm/errors.go at master · go-gorm/gorm](https://github.com/go-gorm/gorm/blob/master/errors.go)  

## Query

[GORM で 1:N で Join したカラムの値を取得する #Go - Qiita](https://qiita.com/zaneli@github/items/f22d57862d780a2a6ca9)  

ネストされた Preload
関連の中にさらに関連を連結したい場合は、Preload("関連名.その先の関連名") のように指定します。

```go
ype User struct {
    ID    uint
    Name  string
    Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
    ID       uint
    UserID   uint
    Amount   float64
    Products []Product `gorm:"foreignKey:OrderID"`
}

type Product struct {
    ID       uint
    OrderID  uint
    Name     string
    Price    float64
}
```

```go
var users []User
err := db.Preload("Orders.Products").Find(&users).Error
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

ネスト先で複数連結する場合は Preload を２回書く
```go
err := db.
    Preload("Orders.Products.Manufacturer").          // Products に関連する Manufacturer
    Preload("Orders.ShippingDetails").                // ShippingDetails
    Find(&users).Error
```

[Scopes | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/ja_JP/docs/scopes.html)  

> Scopes を利用することで、共通で使用されるロジックを再利用することができます。
> 共有ロジックは func(*gorm.DB) *gorm.DB という型として定義します。

[gormのscopesが便利 #Go - Qiita](https://qiita.com/gougyan/items/86f7da4d17a26752a08e)

### Delete

https://gorm.io/ja_JP/docs/delete.html

> gorm.DeletedAt フィールドがモデルに含まれている場合、そのモデルは自動的に論理削除されるようになります。

https://gorm.io/ja_JP/docs/constraints.html



## Customize

[データ型のカスタマイズ | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/ja_JP/docs/data_types.html)  

[GORM Saveメソッドの挙動を追ってみた #Go - Qiita](https://qiita.com/yukiyoshimura/items/3d427ad7168d36a04055)  

> Updateの際、CreatedAtは値指定がなければゼロ値で更新、UpdatedAtは自動生成されるよ
> 
> CreatedAtは他のカラム同様更新対象のカラムとして扱われ、値がなければゼロ値がセットされます。UpdatedAtだけ特別扱いされる形です。
> 
> SaveメソッドでCreatedAtの更新を除外したいときは？
> Omitを使おう


# driver.sqlite

[sqlite package - gorm.io/driver/sqlite - Go Packages](https://pkg.go.dev/gorm.io/driver/sqlite)  

[Connecting to a Database | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/docs/connecting_to_the_database.html#SQLite)  

> NOTE: You can also use file::memory:?cache=shared instead of a path to a file. This will tell SQLite to use a temporary database in system memory. (See SQLite docs for this)

# sqlite3

[sqlite3 package - github.com/mattn/go-sqlite3 - Go Packages](https://pkg.go.dev/github.com/mattn/go-sqlite3)  

## Error Handling

[go-sqlite3/error.go at master · mattn/go-sqlite3](https://github.com/mattn/go-sqlite3/blob/master/error.go)  

errors.As で sqlite3.Error の型を取り出してチェックする必要がる？  
err.Code で 大枠のチェックをして、err.ExtendedCode で制約のチェックをする  

```go
// UNIQUE constraint failed
var sqlErr sqlite3.Error
if errors.As(err, &sqlErr) {
    if sqlErr.Code == sqlite3.ErrConstraint && sqlErr.ExtendedCode == sqlite3.ErrConstraintUnique {
        // Unique 
    }
    if sqlErr.Code == sqlite3.ErrConstraint && sqlErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
        // Unique Foreign Key 
    }
    if sqlErr.Code == sqlite3.ErrConstraint && sqlErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
        // Unique Primary Key
    }
}
```
