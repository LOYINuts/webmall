package middleware

import (
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt中间件，实现token鉴权
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.Success
		// 获取报文头部的authorization
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 过期则返回一个token过期错误
				code = e.ErrorTokenTimeout
			}
		}
		if code != e.Success {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
