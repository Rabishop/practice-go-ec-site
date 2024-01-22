package model

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var err error

func MysqlConn() *gorm.DB {

	Username := "root"
	Password := "kurumi9452"
	Protocol := "tcp"
	Address := "127.0.0.2:3306"
	Dbname := "onlinemarket"

	dsn := Username + ":" + Password + "@" + Protocol + "(" + Address + ")" + "/" + Dbname
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超出的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
	return db
}

var Db *gorm.DB

func init() {
	Db = MysqlConn()
}
