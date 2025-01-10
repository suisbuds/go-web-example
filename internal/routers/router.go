package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/suisbuds/miao/docs"
	"github.com/suisbuds/miao/internal/middleware"
	"github.com/suisbuds/miao/internal/routers/api/v1"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Router 层负责路由注册, 处理 HTTP 请求, 进行参数验证和处理

func NewRouter() *gin.Engine {
	r := gin.New()

	// 中间件注册
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	// 生成 swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 将路由 Handlers 注册到对应的路由组
	article := v1.NewArticle()
	tag := v1.NewTag()

	// 路由接口定义
	apiv1 := r.Group("/api/v1")
	{ 
	
		// 测试 Tag 接口
		// Handlers 不能重复注册: apiv1.GET("/tags", tag.Get)
		apiv1.GET("/tags", tag.List)
		apiv1.POST("/tags", tag.Create)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)
  
		// 测试 Article 接口
		// 验证 Article 接口时必须确保 TagID 和 ArticleID 的对应, 并且根据接口文档调用 (定义好接口后生成接口文档)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
		apiv1.POST("/articles", article.Create)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
