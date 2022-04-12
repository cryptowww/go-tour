package main

import (
	"dbmodel/models"
	"gorm.io/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"fmt"
)

var db *gorm.DB
var err error
var constr string


func main() {
	constr=fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", 3306, "gomicro")

	// 调用数据库迁移
	automigrate()
}

func automigrate(){
	db, err = gorm.Open(mysql.Open(constr),&gorm.Config{})
	if err != nil {
		panic("connect error!")
	}

	db.AutoMigrate(&models.PersonModel{})
}
