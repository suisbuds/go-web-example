package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 针对某些路由接口限流


// 实现 LimiterIface 接口
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

// 根据请求的URI生成限流 Key
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	// 去除 URL 查询参数
	return uri[:index]
}

// 根据 Key 获取对应的 Bucket
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// ...可变传入参数
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}

	return l
}
