package router

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
	"github.com/fastgox/fastgox-api-starter/src/router/middleware"
	"github.com/fastgox/utils/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// HTTP
	AuthRouter   *gin.RouterGroup
	PublicRouter *gin.RouterGroup
	OpenRouter   *gin.RouterGroup
	Engine       *gin.Engine

	// TCP
	TCPRouter *tcp.Router
)

// init 包初始化时创建引擎和路由组
func init() {
	// 初始化 i18n
	i18n.Init()

	Engine = gin.New()
	Engine.Use(gin.Logger(), middleware.RecoveryMiddleware(), middleware.CORSMiddleware(), middleware.I18nMiddleware())

	// 加载HTML模板（支持 layouts 公共模板）
	loadTemplates()

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

	// ========== TCP 路由 ==========
	TCPRouter = tcp.NewRouter()
	TCPRouter.Use(tcp.Recovery(), tcp.Logger())
}

// loadTemplates 加载模板文件，支持 layouts 继承
// 每个页面模板与 layouts/*.html 一起解析，实现公共布局复用
func loadTemplates() {
	layoutFiles, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		logger.Error("加载布局模板失败: %v", err)
	}

	pageFiles, err := filepath.Glob("templates/*.html")
	if err != nil {
		logger.Error("加载页面模板失败: %v", err)
	}

	tmpl := template.New("")
	for _, page := range pageFiles {
		// 每个页面模板与所有 layout 模板一起解析
		files := append([]string{page}, layoutFiles...)
		name := filepath.Base(page)
		t := tmpl.New(name)
		template.Must(t.ParseFiles(files...))
	}

	Engine.SetHTMLTemplate(tmpl)
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
