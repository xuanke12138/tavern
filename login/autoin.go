package login

import (
	"Tavern/db"
	"Tavern/userinterface"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type info struct {
	name string
	password string
}

func check2(err error , c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":-1,
		"err":err,
		"msg":"登陆失败",
	})
}

func AutoIn() gin.HandlerFunc{
	return func (c *gin.Context){

		//获取信息
		tokenn := c.GetHeader("Authorization")
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
			})
			return
		}
		name,_ := c.GetQuery("username")
		password,_ := c.GetQuery("password")


		//fmt.Println(11111111)
		if tokenn != ""{
			claims:=c.MustGet("claims").(*Claim)
			name = claims.Name
			password = claims.Password
		}
		//获取数据库中信息
		var s userinterface.User
		db.Where("user_name = ?",name).Find(&s)


		//验证密码
		err = bcrypt.CompareHashAndPassword([]byte(s.PassWord),[]byte(password))
		if err != nil{
			//check2(err,c)
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"err":err,
				"msg":"密码错误",
			})
			return
		}

		//if s.PassWord != password{
		//	//check2(err,c)
		//	c.JSON(http.StatusOK,gin.H{
		//		"status":-1,
		//		"err":err,
		//		"msg":"333登陆失败",
		//		"y":password,
		//		"k":s.PassWord,
		//	})
		//	return
		//}

		//获取token
		token,err :=CreatToken(name,password)

		if err != nil{
			check2(err,c)
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"status":1,
			"msg":"登陆成功",
			"data":s,
			"token":token,
		})
	}
}