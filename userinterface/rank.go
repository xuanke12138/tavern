package userinterface

import (
	"Tavern/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

func GetByAlcohol() gin.HandlerFunc{
	return func(c *gin.Context){
		//打开数据库
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
		}
		defer db.Close()


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


		var s []User
		db.Model(&s).Order("Alcohol desc").Offset((pagenum-1)*pagesize).Limit(pagesize).Find(&s)

		var summ int
		db.Model(&User{}).Count(&summ)

		temp = c.Query("UserId")
		id,err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}
		alcohol,rank := getrank(id)

		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"date":s,
			"sum":summ,
			"alcohol":alcohol,
			"rank":rank,
		})

	}
}
