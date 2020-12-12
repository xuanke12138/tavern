package comment

import (
	"Tavern/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func FindContent() gin.HandlerFunc{
	return func(c *gin.Context){
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
			return

		}


		word , ok := c.GetQuery("keyword")
		if ok != true || word == ""{
			c.JSON(http.StatusOK,gin.H{
				"msg":"没有请求",
			})
			return
		}

		var s  []Comment
		var whilte Comment
		db.Model(&s).Where("content LIKE ?","%"+word+"%").Order("like_num desc").Find(&s)
		if s[0] == whilte{
			c.JSON(http.StatusOK,gin.H{
				"msg":"无相关信息",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"返回成功",
			"data":s,
		})
	}
}
