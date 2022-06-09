package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Db 全局
var Db *gorm.DB

func init() {
	var err error
	dsn := "root:sunchuan233@tcp(127.0.0.1:3306)/one-admin?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                                   logger.Default.LogMode(logger.Info), //打印sql
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, //禁用创建外键
	})
	if err != nil {
		panic("Connecting database failed:" + err.Error())
	}
}

func GetDB() *gorm.DB {
	return Db
}
