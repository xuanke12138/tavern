package userinterface

import (
	"Tavern/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func GetUserHome() gin.HandlerFunc{
	return func(c *gin.Context){
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
			})
			return
		}
		id := c.Query("UserId")
		var s User
		db.Where("id = ?",id).Find(&s)
		c.JSON(http.StatusOK,gin.H{
			"data":s,
		})
	}
}