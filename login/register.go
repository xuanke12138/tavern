package login

import (
	"Tavern/db"
	"Tavern/userinterface"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"

	//"golang.org/x/crypto/bcrypt"
	"net/http"
)

func check(err error,c *gin.Context){
	c.JSON(-1,gin.H{
		"status": -1,
		"err": err,
		"msg":"注册失败",
	})
}

func Register() gin.HandlerFunc{

	return func (c *gin.Context){
		//获取信息
		db,err := db.GetDb()
		if err != nil{
			check(err,c)
			db.Close()
			return
		}

		db.AutoMigrate(userinterface.User{})



		//加密
		password,_ := c.GetQuery("password")
		hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil{
			check(err,c)
			db.Close()
			return
		}

		encode := string(hash)

		//存储信息
		//result,err:=db.Exec("INSERT INTO student(id,name,age,sex,class,password)VALUES (?,?,?,?,?,?)",id,name,age,sex,class,encodePWD)
		var s , t ,k userinterface.User
		s.UserName , _ = c.GetQuery("username")
		s.PassWord = encode

		db.Where("user_name = ?", s.UserName).Find(&t)
		if t == k{
			db.Create(&s)
			c.JSON(http.StatusOK,gin.H{
				"status":1,
				"msg":"注册成功",
				"data":s,
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg":"用户名已存在",
			})
			return
		}


	}
}

