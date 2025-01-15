package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/suisbuds/miao/docs"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/middleware"
	"github.com/suisbuds/miao/internal/routers/api"
	v1 "github.com/suisbuds/miao/internal/routers/api/v1"
	"github.com/suisbuds/miao/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router 层负责注册路由, 调用 API 端点, 利用对应 Handlers 处理 HTTP 请求, 并进行参数校验

// 针对 auth 路由接口限流
var authLimiter = limiter.NewRouteLimiter().AddBuckets(
	// 令牌桶规则
	limiter.LimiterBucketRule{
		Key:          "/auth",     // 路由 URI
		FillInterval: time.Second, // 每秒生成一个令牌
		Capacity:     10,          // 容量
		Quantum:      1,           // 每次填充一个令牌
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 中间件注册
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
		r.Use(middleware.AccessLog()) // 访问日志
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery()) // 捕获 panic 并发送邮件预警
	}
	r.Use(middleware.RateLimiter(authLimiter)) // 限流器
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout)) // 超时控制
	r.Use(middleware.Translations())          // 翻译错误信息
	r.Use(middleware.AppInfo())               // 获取应用信息


	// 注册 API 端点
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()

	// 注册路由接口请求并配置相应 Handlers
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册 GET 请求 swagger 文档路由
	r.POST("/upload/file", upload.UploadFile)                            // 注册 POST 请求文件上传路由
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))    // 设置静态文件服务: 将 /static 路径映射到服务器的 UploadSavePath 目录, 然后通过访问 /static URL 获取上传的静态文件
	r.POST("/auth", api.GetAuth)                                         // 注册 POST 请求认证路由

	// 注册路由组, 定义路由组接口
	apiv1 := r.Group("/api/v1")

	//  分组路由中间件注册
	if global.ServerSetting.RunMode == "debug" {
		// 针对 apiv1 路由组使用 JWT
		apiv1.Use(middleware.JWT())
	}

	{
		// 注册 Tag & Article 路由接口并配置 Handlers

		// 测试 Tag 接口: 创建 Tag, 获取 Tag 列表, 更新 Tag, 删除 Tag
		apiv1.POST("/tags", tag.Create)
		apiv1.GET("/tags", tag.List)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)

		// 验证 Article 接口时必须确保 TagID 和 ArticleID 的对应, 并且根据接口文档调用 (定义好接口后再生成接口文档)
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
