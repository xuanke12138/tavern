package main

import (
	"Tavern/comment"
	"Tavern/db"
	"Tavern/login"
	"Tavern/userinterface"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	e := gin.New()
	e.Use(gin.Logger(),gin.Recovery())

	//打开数据库，只开一次节约时间
	db := db.Open()
	defer db.Close()

	e.GET("/register",login.Register())
	e.GET("/autoin",login.VareToken(),login.AutoIn())
	e.GET("/in",login.In())

	e.GET("/upcomment",comment.AddComment())
	e.GET("/upreply",comment.AddReply())
	e.GET("/zan",comment.CreatLike())

	e.GET("/deletecomment",comment.DeleteComment())
	e.GET("/deletereply",comment.DeleteReply())


	e.GET("/getcommentbyzan",comment.GetCommentByZan())
	e.GET("/getrank",userinterface.GetByAlcohol())
	e.GET("/getcommentbytime",comment.GetCommentByTime())
	e.GET("/getreply",comment.GetReply())
	e.GET("/getcarecomment",comment.GetCommentFriend())

	e.GET("/findcomment",comment.FindContent())
	e.GET("/getcommentrand",comment.GetCommentRand())

	e.GET("/getuserhome",userinterface.GetUserHome())
	e.GET("/addcare",userinterface.AddCare())
	e.GET("/deletecare",userinterface.DeleteCare())
	e.GET("/getcare",userinterface.GetFriend())

	//comment.Po(c)
	e.Run(":8080")


}
