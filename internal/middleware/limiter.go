package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/limiter"
)



func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		// 放行
		if bucket, ok := l.GetBucket(key); ok {
			// 尝试获取令牌
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				// 令牌超限, 控制接口请求速度
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
