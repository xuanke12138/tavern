package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VareToken() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenheader := c.GetHeader("Authorization")

		if tokenheader ==""{
			c.Next()
			c.Abort()
			return
		}


		token := strings.Split(tokenheader, " ")

		claims , err:= ParseToken(token[0])


		//log.Println(claims.Password,err)
		if err != "nil"{
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"err":err,
				"msg":"token错误",
				"token":token,
			})
			c.Abort()
			return
		}

		c.Set("claims",claims)
		c.Next()
	}
}
