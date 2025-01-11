package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/suisbuds/miao/docs"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/middleware"
	"github.com/suisbuds/miao/internal/routers/api"
	"github.com/suisbuds/miao/internal/routers/api/v1"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Router 层负责路由注册, 处理 HTTP 请求, 调用 API 端点, 进行参数验证和处理

func NewRouter() *gin.Engine {
	r := gin.New()


	// 中间件注册
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations()) // 错误信息翻译
	

	// 将路由 Handlers 注册到路由组
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册 GET 请求 swagger 文档路由
	r.POST("/upload/file", upload.UploadFile) // 注册 POST 请求文件上传路由
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) // 设置静态文件服务: 将 /static 路径映射到服务器的 UploadSavePath 目录, 然后通过访问 /static URL 获取上传的静态文件


	// 路由接口定义
	apiv1 := r.Group("/api/v1")
	{ 
	
		// Handlers 不能重复注册: apiv1.GET("/tags", tag.Get)
		// 测试 Tag 接口: 创建 Tag, 获取 Tag 列表, 更新 Tag, 删除 Tag
		apiv1.POST("/tags", tag.Create)
		apiv1.GET("/tags", tag.List)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)
  
		// 验证 Article 接口时必须确保 TagID 和 ArticleID 的对应, 并且根据接口文档调用 (定义好接口后生成接口文档)
		// 测试 Article 接口: 创建 Article, 获取指定 Article, 获取 Article 列表, 更新 Article, 删除 Article
		apiv1.POST("/articles", article.Create)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
