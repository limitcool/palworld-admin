package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/limitcool/starter/pkg/code"
)

// Cors 跨域解决
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

// AdminPasswordMiddleware 是用于验证管理员密码的中间件
func AdminPasswordMiddleware(expectedPassword string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取密码
		password := c.GetHeader("Password")

		// 验证密码是否为管理员密码
		if password != expectedPassword {
			code.AutoResponse(c, nil, code.NewErrCode(code.UserAuthFailed))
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next()
	}
}
