package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 对应用接口进行限流, 防止恶意请求和服务器过载
// ratelimit 提供令牌桶算法, 简化实现限流器
// 空接口类型变量可以接收任意类型的变量, 通过类型断言获取具体类型

// 限流器接口, 可以根据应用景实现不同的限流器
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// 限流器
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket // 令牌桶
}

// 令牌桶配置
type LimiterBucketRule struct {
	Key          string // ID
	FillInterval time.Duration // 令牌填充间隔
	Capacity     int64 // 令牌桶容量
	Quantum      int64 // 每次生成令牌数量
}
