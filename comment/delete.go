package comment

import (
	"Tavern/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

func DeleteComment() gin.HandlerFunc{
	return func (c *gin.Context){
		//打开数据库
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
			return
		}


		temp := c.Query("deleteid")
		value2, err := strconv.Atoi(temp)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		id := value2

		//删除评论
		db.Where("id = ?",id).Delete(&Comment{})

		c.JSON(http.StatusOK,gin.H{
			"msg":"删除成功",
		})
	}
}

func DeleteReply() gin.HandlerFunc{
	return func (c *gin.Context){
		//打开数据库
		db,err := db.GetDb()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
			return
		}


		temp := c.Query("deleteid")
		value2, err := strconv.Atoi(temp)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		id := value2

		//删除评论
		db.Where("id = ?",id).Delete(&Reply{})

		c.JSON(http.StatusOK,gin.H{
			"msg":"删除成功",
		})
	}
}
