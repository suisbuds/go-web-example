package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/suisbuds/miao/docs"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/middleware"
	"github.com/suisbuds/miao/internal/router/api"
	v1 "github.com/suisbuds/miao/internal/router/api/v1"
	"github.com/suisbuds/miao/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router 层负责注册路由, 调用 API Controller

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
	r.Use(middleware.Cors())                                           // 跨域请求
	r.Use(middleware.RateLimiter(authLimiter))                         // 限流器
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout)) // 超时控制
	r.Use(middleware.Translations())                                   // 翻译错误信息
	r.Use(middleware.AppInfo())                                        // 获取应用信息

	// 注册 api 路由端点
	user := v1.NewUser()
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	jwt := middleware.InitJWT()

	// 注册 api 路由接口
	r.POST("/login", jwt.LoginHandler)                                   // 注册 POST 请求登录路由
	r.GET("/refresh", jwt.RefreshHandler)                                // 注册 GET 请求刷新令牌路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册 GET 请求 swagger 文档路由
	r.POST("/upload/file", upload.UploadFile)                            // 注册 POST 请求文件上传路由
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))    // 设置静态文件服务: 将 /static 路径映射到服务器的 UploadSavePath 目录, 然后通过访问 /static URL 获取上传的静态文件
	r.POST("/auth", api.GetAuth)                                         // 注册 POST 请求认证路由

	// 注册路由组, 分组注册中间件
	apiAuth := r.Group("/api/auth")
	apiAuth.Use(jwt.MiddlewareFunc()) // JWT 中间件

	apiv1 := r.Group("/api/v1")
	{
		// 根据接口注解调用

		// 注册 Tag 接口
		apiv1.POST("/tags", tag.CreateTag)
		apiv1.GET("/tags", tag.GetTagList)
		apiv1.PUT("/tags/:id", tag.UpdateTag)
		apiv1.PATCH("/tags/:id/state", tag.UpdateTag)
		apiv1.DELETE("/tags/:id", tag.DeleteTag)

		// 注册 Article 接口
		apiv1.POST("/articles", article.CreateArticle)
		apiv1.GET("/articles/:id", article.GetArticle)
		apiv1.GET("/articles", article.GetArticleList)
		apiv1.PUT("/articles/:id", article.UpdateArticle)
		apiv1.PATCH("/articles/:id/state", article.UpdateArticle)
		apiv1.DELETE("/articles/:id", article.DeleteArticle)

		// 注册 User 接口
		apiv1.POST("/user", user.CreateUser)
		apiv1.GET("/user/:username", user.GetUser)
		apiv1.GET("/user", user.GetUserList)
		apiv1.PUT("/user/:username", user.UpdateUser)
		apiv1.DELETE("/user/:username", user.DeleteUser)
	}

	return r
}
