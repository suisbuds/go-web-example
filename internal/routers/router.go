package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/suisbuds/miao/internal/routers/api/v1"
	_ "github.com/suisbuds/miao/docs"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	// 测试接口
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		// apiv1.DELETE("/tags/:id", tag.Delete)

		apiv1.GET("/articles/:id", article.Get)
		// apiv1.POST("/articles", article.Create)
		// apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
