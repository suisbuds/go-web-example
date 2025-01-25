package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许访问所有域
		c.Header("Access-Control-Allow-Credentials", "true")      // 跨域请求携带 cookie
		//  header 支持类型
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 支持的跨域请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		c.Header("Access-Control-Max-Age", "21600") // 缓存结果超时设定
		c.Set("content-type", "application/json")   // 返回格式
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
