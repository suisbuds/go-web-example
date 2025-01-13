package middleware

import "github.com/gin-gonic/gin"

// 获取应用内部信息和业务信息
// gin.Context 的 setter&getter (元数据管理), 对 interface 进行类型断言

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "miao")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
