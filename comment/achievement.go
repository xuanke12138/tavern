package comment

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const(
	第一次评论 = "xxxx"
)

type Achievement struct {
	FromId int `gorm:"NOT NULL"`
	Name string `gorm:"varchar(16)"`
	Content string`gorm:"varchar(32)"`
	gorm.Model
}

func AchieveComment(db *gorm.DB,id int) (bool,Achievement){
	var summ int
	db.Model(&Comment{}).Where("from_id = ?",id).Count(&summ)
	switch summ  {
	case 1:
		var s Achievement
		s.FromId = id
		s.Name = 第一次评论
		db.Create(&s)
		return true,s
	default:
		return false,Achievement{}
	}
}
