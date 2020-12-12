package login

import (
	"Tavern/db"
	"Tavern/userinterface"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func In() gin.HandlerFunc{
	return func(c *gin.Context){
		//获取信息
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
			})
			return
		}
		name ,_ := c.GetQuery("username")
		password , _ := c.GetQuery("password")
		//获取数据库中信息
		var s userinterface.User
		db.Where("user_name = ?",name).Find(&s)

		//验证密码
		err = bcrypt.CompareHashAndPassword([]byte(s.PassWord),[]byte(password))
		if err != nil{
			check2(err,c)
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status":1,
			"msg":"登陆成功",
			"data":s,
		})
	}
}