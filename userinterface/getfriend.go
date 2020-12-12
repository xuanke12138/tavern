package userinterface

import (
	"Tavern/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type Care struct {
	ToUserId int `gorm:"NOT NULL"`
	ToUserName string `gorm:"varchar(12)"`
	FromUserId int `gorm:"NOT NULL"`
	FromUserName string `gorm:"varchar(12)"`
	gorm.Model
}

func AddCare() gin.HandlerFunc{
	return func(c *gin.Context){
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
			})
			return
		}
		db.AutoMigrate(&Care{})
		var s Care
		value := c.Query("FromId")
		//因为query的是字符串所以要转化为数字
		value2, err := strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		s.FromUserId = value2
		s.FromUserName = c.Query("FromName")

		value = c.Query("ToId")
		//因为query的是字符串所以要转化为数字
		value2, err = strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		s.ToUserId = value2
		s.ToUserName = c.Query("ToName")

		var s1,s2 Care
		db.Where("from_user_id = ? AND to_user_id = ?",s.FromUserId,s.ToUserId).Find(&s1)
		if s1 != s2{
			c.JSON(http.StatusOK,gin.H{
				"msg":"已关注",
			})
			return
		}
		db.Create(&s)
		c.JSON(http.StatusOK,gin.H{
			"msg":"关注成功",
			"data":s,
		})
	}
}
func DeleteCare() gin.HandlerFunc{
	return func(c *gin.Context){
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
			})
			return
		}
		db.AutoMigrate(&Care{})
		var s Care
		value := c.Query("FromId")
		//因为query的是字符串所以要转化为数字
		value2, err := strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		s.FromUserId = value2
		s.FromUserName = c.Query("FromName")

		value = c.Query("ToId")
		//因为query的是字符串所以要转化为数字
		value2, err = strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		s.ToUserId = value2
		s.ToUserName = c.Query("ToName")

		var s1,s2 Care
		db.Where("from_user_id = ? AND to_user_id = ?",s.FromUserId,s.ToUserId).Find(&s1)
		if s1 == s2{
			c.JSON(http.StatusOK,gin.H{
				"msg":"未关注",
			})
			return
		}
		db.Where("from_user_id = ? AND to_user_id = ?",s.FromUserId,s.ToUserId).Delete(&s)
		c.JSON(http.StatusOK,gin.H{
			"msg":"删除关注成功",
			"data":s,
		})
	}
}


func GetFriend() gin.HandlerFunc{
	return func(c *gin.Context){
		//打开数据库
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
			return
		}


		var care []Care

		value := c.Query("FromId")

		value2, err := strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		userid := value2

		//按页数获得信息
		temp := c.Query("PageSize")
		pagesize,err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}

		temp = c.Query("PageNum")
		pagenum,err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}

		//获取关注的人
		var sum int
		db.Model(&care).Where("from_user_id = ?",userid).Count(&sum)
		db.Model(&care).Where("from_user_id = ?",userid).Offset((pagenum-1)*pagesize).Limit(pagesize).Find(&care)

		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"data":care,
			"sum":sum,
		})
	}
}


