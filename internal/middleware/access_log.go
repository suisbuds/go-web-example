package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/logger"
)

// 错误日志, 业务日志, 访问日志
// 访问日志: 请求方法, 状态码, 开始时间, 结束时间, 请求参数和响应数据记录到访问日志, 实现日志链路追踪

type AccessLogWriter struct {
	gin.ResponseWriter // 嵌入
	body *bytes.Buffer // 存储响应字段
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter // 拦截响应数据

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		str := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"

		global.Accesser.WithFields(fields).Logf(logger.INFO, logger.SINGLE, str,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
