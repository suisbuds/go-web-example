package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/email"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/logger"
)

// 捕获 panic 并发送邮件预警

func Recovery() gin.HandlerFunc {
	// 邮件发送器
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	// 闭包函数
	return func(c *gin.Context) {
		defer func() {
			// recover 捕获 panic, 防止程序崩溃
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Logf(logger.ERROR, logger.SINGLE, "panic recover err: %v", err)

				// 邮件预警
				err := mailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("PANIC happens in: %d", time.Now().Unix()),
					fmt.Sprintf("ERROR message: %v", err),
				)
				if err != nil {
					global.Logger.Logf(logger.PANIC, logger.FRAMES, "mail.SendMail err: %v", err)
				}

				// 返回错误响应
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort() // 终止请求
			}
		}()
		c.Next() // 继续执行请求
	}
}
