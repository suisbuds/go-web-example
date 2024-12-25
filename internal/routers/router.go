package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/suisbuds/miao/docs"
	"github.com/suisbuds/miao/internal/middleware"
	v1 "github.com/suisbuds/miao/internal/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 中间件注册
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 编写路由 Handlers，注册到路由对应的规程
	article := v1.NewArticle()
	tag := v1.NewTag()

	// 测试接口
	apiv1 := r.Group("/api/v1")
	{
		
		// apiv1.GET("/tags", tag.Get) // Handlers 不能重复注册
		apiv1.GET("/tags", tag.List)
		apiv1.POST("/tags", tag.Create)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)
  
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
		apiv1.POST("/articles", article.Create)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
