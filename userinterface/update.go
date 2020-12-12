package userinterface

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct{
	Alcohol int `gorm:"default 10"`
	UserName string `gorm:"varchar(16);unique`
	PassWord string `gorm:"varchar(32)"`
	Money int `gorm:"default 0"`
	AchievementNum int `gorm:"default 0"`
	gorm.Model
}
