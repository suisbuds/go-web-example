package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/pkg/errcode"
)

// 404
func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    errcode.NotFound.Code,
		"message": errcode.NotFound.Msg,
	})
}
