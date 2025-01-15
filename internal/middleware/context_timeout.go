package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/pkg/errcode"
)

// 统一超时控制, 避免上下游应用互相影响, 造成连锁反应

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// 创建携带超时的 context, 使用该 context 的服务会进行超时控制
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel() // 防止 context 泄漏

		// 将超时 context 注入服务
		c.Request = c.Request.WithContext(ctx)

		doneChan := make(chan struct{})     // 请求正常完成
		panicChan := make(chan interface{}) // 捕获 panic

		// goroutine 处理请求
		go func() {
			// 捕获 panic 交由 panicChan 处理
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			c.Next()        // 执行后续中间件
			close(doneChan) // 关闭 doneChan
		}()

		// 等待三种事件之一发生
		select {
		case p := <-panicChan:
			// 重新抛出 panic, 交由 Recovery 中间件处理
			panic(p)
		case <-doneChan:
			// doneChan 关闭, 请求正常完成
		case <-ctx.Done():
			// 请求超时
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"code": errcode.RequestTimeout.Code(),
				"msg":  errcode.RequestTimeout.Msg(),
			})
		}
	}
}
