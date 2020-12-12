package comment

import (
	"Tavern/db"
	"Tavern/userinterface"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type Comment struct {
	Content string `gorm:"varchar(512);not null"`
	LikeNum int `gorm:"default 0"`
	ReplyNum int `gorm:"default 0"`
	FromId int `gorm:"not null"`
	FromName string `gorm:"varchar(16);not null"`
	Topic string `gorm:"varchar(32);not mull"`

	gorm.Model
}

type Reply struct {
	ToCommentId int `gorm:"not null"`
	Content string `gorm:"varchar(512);not null"`
	ToId int `gorm:"not null"`
	LikeNum int `gorm:"default 0"`
	FromId int `gorm:"not null"`
	FromName string `gorm:"varchar(16);not null"`
	gorm.Model
}

type Like struct {
	ToId int `gorm:"not null"`
	FromId int `gorm:"not null"`
	gorm.Model
}

func AddComment() gin.HandlerFunc {
	return func (c * gin.Context){
		//打开数据库
		db,err := db.GetDb()
		db.LogMode(true)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "打开数据库失败",
				"err":err,
			})
			return
		}


		//获取评论数据
		var s Comment
		s.Content = c.Query("Content")
		value := c.Query("FromId")
		//因为query的是字符串所以要转化为数字
		value2, err := strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		s.FromId = value2
		s.FromName = c.Query("FromName")
		s.Topic = c.Query("Topic")
		//s.Time = time.Now()

		//根据Comment创建表单，已存在无影响
		//db.Table("xxxx").Create(&Conments{})创建指定名称表单，但若是已存在则有问题，会报错
		db.AutoMigrate(&Comment{})

		//插入评论
		db.Create(&s)

		//查看成就
		if tr,_ := AchieveComment(db,s.FromId); tr{
			c.JSON(http.StatusOK, gin.H{
				"msg": "插入评论成功",
				"data":s,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"msg": "插入评论成功",
				"data":s,
			})
		}


		//获取评论的酒精值
		var user userinterface.User
		db.Where("id = ?",s.FromId).Find(&user)
		db.Model(&user).Where("id = ?",s.FromId).Update(map[string]interface{}{"Alcohol":user.Alcohol+10})

	}
}



func AddReply() gin.HandlerFunc {
	return func (c * gin.Context){

	//打开数据库
	db,err := db.GetDb()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
		"msg": "打开数据库失败",
		})
		return
	}


	//获取评论数据
	var s Reply
	s.Content = c.Query("Content")
	s.FromName = c.Query("FromName")

	value := c.Query("FromId")
	value2, err := strconv.Atoi(value)
	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
		"msg": "转化id失败",
		})
		return
	}
	s.FromId = value2

	value = c.Query("ToId")
	value2, err = strconv.Atoi(value)
	if err != nil && value != ""{
		c.JSON(http.StatusOK, gin.H{
			"msg": "转化id失败",
		})
		return
	}
	s.ToId = value2

	value = c.Query("ToCommentId")
	value2, err = strconv.Atoi(value)
	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "转化id失败",
		})
		return
	}
	s.ToCommentId = value2


	db.AutoMigrate(&Reply{})

	//插入评论
	db.Create(&s)
	c.JSON(http.StatusOK, gin.H{
		"msg": "插入回复成功",
		"data":s,
		})


	//获取回复的酒精值
	var user,user2 userinterface.User
	var com Comment

	db.Model(&com).Where("id = ?",s.ToCommentId).Find(&com)
	fmt.Println(com.ReplyNum,com.ID,5555,com)
	db.Model(&com).Where("id = ?",s.ToCommentId).Update(map[string]interface{}{"reply_num":com.ReplyNum+1})

	db.Where("id = ?",s.FromId).Find(&user)
	db.Model(&user).Where("id = ?",s.FromId).Update(map[string]interface{}{"Alcohol":user.Alcohol+5})

	//获取收到回复酒精值
	db.Where("id = ?",com.FromId).Find(&user2)
	db.Model(&user2).Where("id = ?",com.FromId).Update(map[string]interface{}{"Alcohol":user2.Alcohol+5})
	}
}


func CreatLike() gin.HandlerFunc{
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



		//获取变量
		var like Like
		temp := c.Query("ToId")
		like.ToId,err = strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}
		temp = c.Query("FromId")
		like.FromId,err = strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}
		db.AutoMigrate(&Like{})

		//搜索是否已经赞过
		//AND 查询时的并列条件
		var w , kong Like
		db.Where("to_id = ? AND from_id = ?",like.ToId,like.FromId).Find(&w)
		if w != kong{
			c.JSON(http.StatusOK,gin.H{
				"msg":"别点了，哥!",
			})
			return
		}

		//给评论加赞
		var s Comment
		var user userinterface.User

		db.Where("id = ?",like.ToId).Find(&s)
		db.Model(&s).Where("id = ?",like.ToId).Update(map[string]interface{}{"like_num":s.LikeNum+1})
		c.JSON(http.StatusOK,gin.H{
			"msg":"赞成功",
			"data":s,
		})

		//赞人酒精值
		db.Model(&user).Where("id = ?",like.FromId).Find(&user)
		db.Model(&user).Where("id = ?",like.FromId).Update(map[string]interface{}{"alcohol":user.Alcohol+5})

		//将赞加入赞表
		db.Create(&like)

		//获取得赞的酒精值

		db.Where("id = ?",s.FromId).Find(&user)
		db.Model(&user).Where("id = ?",s.FromId).Update(map[string]interface{}{"Alcohol":user.Alcohol+5})


	}
}

