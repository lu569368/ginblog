package common

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormInit() *gorm.DB {
	var err error
	var Db *gorm.DB
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	return Db
}
