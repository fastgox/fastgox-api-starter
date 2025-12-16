package router

import (
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	AuthRouter   *gin.RouterGroup
	PublicRouter *gin.RouterGroup
	OpenRouter   *gin.RouterGroup
	Engine       *gin.Engine
)

// init 包初始化时创建引擎和路由组
func init() {
	Engine = gin.Default() // 使用Default()自动包含Logger和Recovery中间件
	Engine.Use(middleware.CORSMiddleware())

	// 加载HTML模板
	Engine.LoadHTMLGlob("templates/*")

	frontPrefix := "/api/v1"
	PublicRouter = Engine.Group(frontPrefix)
	AuthRouter = Engine.Group(frontPrefix)
	AuthRouter.Use(middleware.AuthMiddleware())

	// 静态文件服务
	setupStaticFiles()

	// 设置模板路由
	setupTemplateRoutes()

	// Swagger 文档路由
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// setupTemplateRoutes 设置模板路由
func setupTemplateRoutes() {
	// 首页
	Engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "FastGoX API Starter",
			"message": "欢迎使用 FastGoX API 服务",
			"version": "1.0.0",
		})
	})

	// 404 页面
	Engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"code":    404,
			"title":   "页面未找到",
			"message": "您访问的页面不存在",
		})
	})
}

// setupStaticFiles 设置静态文件服务
func setupStaticFiles() {
	if config.GlobalConfig != nil {
		fileConfig := config.GlobalConfig.File
		// 提供静态文件访问服务
		Engine.Static(fileConfig.URLPrefix, fileConfig.UploadPath)
	}
}
