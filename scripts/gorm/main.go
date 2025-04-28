package main

import (
	"time"
	"fmt"
	"database/sql"
	"database/sql/driver"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

//	"github.com/jinzhu/now"
)

type OrgField sql.NullString

// Scan implements the Scanner interface.
func (n *OrgField) Scan(value interface{}) error {
	fmt.Println("--- Scan")
	fmt.Println(value)
	return (*sql.NullString)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n OrgField) Value() (driver.Value, error) {
	fmt.Println("--- Value")
	fmt.Println(n.Valid)
	if !n.Valid {
		return nil, nil
	}
	fmt.Println(n.Value)
	return n.Value, nil
}

func (OrgField) QueryClauses(f *schema.Field) []clause.Interface {
	fmt.Println("--- QueryClauses")
	fmt.Println(f)
	fmt.Println(f.TagSettings)
	return []clause.Interface{gorm.SoftDeleteQueryClause{Field: f, ZeroValue: parseZeroValueTag(f)}}
}

func parseZeroValueTag(f *schema.Field) sql.NullString {
/*
	return sql.NullString{String: v, Valid: true}
	if v, ok := f.TagSettings["ZEROVALUE"]; ok {
		if _, err := now.Parse(v); err == nil {
		}
	}
*/
	return sql.NullString{Valid: false}
}

//----

type User struct {
	ID        string `gorm:"primarykey;size:16"`
	Name      string `gorm:"index"`
//	Key       OrgField `gorm:"hoge:fuga"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Organization struct {
	ID        string `gorm:"primarykey;size:16"`
	Name      string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationUser struct {
	ID             string `gorm:"primarykey;size:16"`
	OrganizationID string
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	UserID         string
	User           User  `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
	  panic("failed to connect database")
	}

	// 外部キー制約をONにする
	db.Exec("PRAGMA foreign_keys = ON;")

	db.AutoMigrate(
		&User{},
		&Organization{},
		&OrganizationUser{},
	)

//	db.Create(&User{ID: "0001", Name: "太郎", Key: OrgField{String: "search1", Valid: true}})
//	db.Save(&User{ID: "0002", Name: "二郎", Key: OrgField{String: "search2", Valid: true}})
	fmt.Println("--- Create ---")
	db.Create(&User{ID: "0001", Name: "太郎"})

	fmt.Println("--- Save ---")
	db.Omit("CreatedAt").Save(&User{ID: "0002", Name: "二郎"})

	u1 := &User{ID: "0002"}
	if err := db.First(u1).Error; err != nil {
		panic(err.Error())
	}

	fmt.Println("---- Search Pattern 1 ----")
	fmt.Println(u1)
/*
	u2 := &User{Key: OrgField{String: "search2", Valid: true}}
	if err := db.First(u2).Error; err != nil {
		panic(err.Error())
	}

	fmt.Println("---- Search Pattern 2 ----")
	fmt.Println(u2)
*/
	/*
	u3 := &User{}
	if err := db.First(u3, User{Name: "二郎"}).Error; err != nil {
		panic(err.Error())
	}

	fmt.Println("---- Search Pattern 3 ----")
	fmt.Println(u3)
*/

	// 外部キー
	db.Create(&User{ID: "1001", Name: "山田"})
	db.Create(&Organization{ID: "0001", Name: "ECShop"})
	db.Create(&OrganizationUser{ID: "0001", OrganizationID: "0001", UserID: "1001"})

	db.Delete(&User{ID: "1001"})

}
