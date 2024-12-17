package app

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/convert"
)

// 对接口返回的结果进行分页处理

func GetPage(c *gin.Context) int {
	page := convert.ConvertStr(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.ConvertStr(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
