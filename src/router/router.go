package router

import (
	"github.com/fastgox/fastgox-api-starter/src/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	ginSwaggerFiles "github.com/swaggo/files"
)

var (
	AuthRouter   *gin.RouterGroup
	PublicRouter *gin.RouterGroup
	Engine       *gin.Engine
)

// init 包初始化时创建引擎和路由组
func init() {
	Engine = gin.New()
	Engine.Use(middleware.CORSMiddleware())
	// Swagger 文档路由
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	frontPrefix := "/api/v1"
	PublicRouter = Engine.Group(frontPrefix)
	AuthRouter = Engine.Group(frontPrefix)
	AuthRouter.Use(middleware.AuthMiddleware())
}
