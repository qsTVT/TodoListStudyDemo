package middleware

import (
	"github.com/gin-gonic/gin"
	"golang/TodoList/pkg/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			msg = "请先登陆"
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 401 //token是假的
				msg = "Token是假的"
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 403 //token无效
				msg = "Token过期"
			}
		}
		if code != 200 {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
