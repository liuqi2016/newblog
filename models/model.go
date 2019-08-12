package models

import (
	"blog/utils/setting"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	//初始化mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	err                                               error
	dbType, dbName, user, password, host, tablePrefix string
	db                                                *gorm.DB
)

func NewDB() (db *gorm.DB) {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	//数据库连接
	db, err = gorm.Open(dbType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName,
		))
	if err != nil {
		panic(err)
	}
	//数据库前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
	// 自动迁移模式
	db.AutoMigrate(&Users{})
	return
}

// func NewDB() (db *gorm.DB) {
// 	return db
// }

//CloseDB 关闭数据库
func CloseDB() {
	defer db.Close()
}

//ID 主键
type ID struct {
	ID int64 `json:"id"`
}
