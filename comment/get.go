package comment

import (
	"Tavern/db"
	"Tavern/userinterface"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)


func GetCommentByTime() gin.HandlerFunc{
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

		//从数据库获取一条id为1的数据并绑定到s上,
		//var s Comment
		//db.(where id = 1).Find(&S)

		//从数据库获得多条数据并绑定到s上,按照时间降序排序后
		//不加des是默认升序
		//order("xxx")按xxx排序
		var s []Comment
		db.Order("created_at desc").Offset(pagesize*pagenum-pagesize).Limit(pagesize).Find(&s)

		//获取评论总条数
		var summ int
		db.Model(Comment{}).Count(&summ)

		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"date":s,
			"sum":summ,
		})

	}
}

func GetCommentByZan() gin.HandlerFunc{
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



		//实现分页功能
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


		var s []Comment
		db.Order("like_num desc").Offset((pagenum-1)*pagesize).Limit(pagesize).Find(&s)

		var summ int
		db.Model(Comment{}).Count(&summ)

		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"date":s,
			"sum":summ,
		})

	}
}

func GetReply() gin.HandlerFunc{
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


		temp := c.Query("CommentId")
		id,err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"转化失败",
				"err":err,
			})
			return
		}

		temp = c.Query("PageSize")
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

		var s []Reply
		db.Where("to_comment_id = ?",id).Order("created_at").Offset((pagenum-1)*pagesize).Limit(pagesize).Find(&s)

		var summ int
		db.Model(Reply{}).Where("to_comment_id = ?",id).Count(&summ)

		var comment Comment
		db.Where("id = ?",id).Find(&comment)

		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"comment":comment,
			"date":s,
			"sum":summ,
		})

	}
}


func GetCommentRand() gin.HandlerFunc{
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


		var s ,s1[]Comment
		var ans []Comment
		var flag [15]int

		db.Model(&Comment{}).Order("created_at desc").Limit(2000).Find(&s1).Order("like_num+reply_num*5 desc",true).Find(&s)
		var siz float64
		db.Model(&Comment{}).Limit(2000).Count(&siz)

		maxx1 := math.Min(2000,siz)
		maxx := int(maxx1)
		maxn := 0

		for i := 0;i < maxx;i++{
			maxn = maxn + s[i].LikeNum + (s[i].ReplyNum*5)
		}



		summ := 0
		k := 0
		for k = -1; k<int(math.Min(9,maxx1))-1;{
			//生成随机数
			rand.Seed(time.Now().UnixNano()*571)
			num := rand.Intn(maxn)
			for i := 0;i < maxx;i++{
				summ+=s[i].LikeNum+(s[i].ReplyNum*5)
				if summ > num && flag[i] != 999{
					k++
					ans = append(ans, s[i])
					flag[i]=999
					break
				}
			}
		}

		c.JSON(http.StatusOK,gin.H{
			"msg":"刷新成功",
			"data":ans,
		})



	}
}


type Comments []Comment

//Len()
func (s Comments) Len() int {
	return len(s)
}

//Less():时间将有高到低排序
func (s Comments) Less(i, j int) bool {
	//时间比大小要转为unix形式，另外unix是in64形式
	return s[i].CreatedAt.Unix() > s[j].CreatedAt.Unix()
}

//Swap()
func (s Comments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func GetCommentFriend() gin.HandlerFunc{
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


		var s,s1 Comments
		var care []userinterface.Care
		var siz int

		value := c.Query("FromId")

		value2, err := strconv.Atoi(value)
		if (err != nil) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "转化id失败",
			})
			return
		}
		userid := value2

		//获取关注的人
		db.Model(&care).Where("from_user_id = ?",userid).Find(&care).Count(&siz)

		//根据关注的人获得评论
		for i:=0;i<siz;i++{
			db.Model(&s).Where("from_id = ?",care[i].ToUserId).Scan(&s1)
			s = append(s, s1...)
		}

		//将关注评论按时间排序
		sort.Sort(s)
		sum := len(s)

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

		var ss []Comment
		for i:=(pagenum-1)*pagesize;i<pagenum*pagesize;i++{

			if i >= sum{
				break
			}
			ss = append(ss, s[i])
		}


		c.JSON(http.StatusOK,gin.H{
			"msg":"获取成功",
			"data":ss,
			"sum":sum,
		})
	}
}

