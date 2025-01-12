package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	// 闭包: 闭包是一个函数, 允许访问其外部作用域的变量, 即使外部函数作用域结束, 闭包仍可以访问变量, 因此可能导致内存泄漏
	// 创建闭包, 在中间件初始化中注入依赖, 传递参数, 提高中间件的可扩展性
	return func(c *gin.Context) {
		
		ecode := errcode.Success
		
		// 从查询参数或请求头中获取 token
		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				// 错误类型判断
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					// 过期 token
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}

		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return

		}

		c.Next()
	}
}
