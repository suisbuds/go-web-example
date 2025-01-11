package api

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/service"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/convert"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/logger"
	"github.com/suisbuds/miao/pkg/upload"
	"go.uber.org/zap/zapcore"
)

// 处理文件上传的 API 端点

type Upload struct{}

func (u Upload) GetSavePath() {
	panic("unimplemented")
}

func NewUpload() Upload {
	return Upload{} // 返回结构体模板的实例
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 请求文件和文件头
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	// 检查文件头和文件类型有效性
	fileType := convert.ConvertStr(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	// 调用文件上传服务
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.UploadFile err: %v", err)
		global.Zapper.Logf(zapcore.ErrorLevel, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	// 文件访问 URL
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
