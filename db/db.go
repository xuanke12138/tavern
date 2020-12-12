package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Open() *gorm.DB{
	//var err error
	db , err = gorm.Open("mysql", "root:20020112a@(127.0.0.1)/tavern?charset=utf8mb4&loc=Local&parseTime=true")
	//fmt.Println(err,5555)
	return db
}

func GetDb() (*gorm.DB,error){
	return db,err
}
