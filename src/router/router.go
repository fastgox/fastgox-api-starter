package router

import (
	"github.com/fastgox/fastgox-api-starter/src/router/middleware"
	"github.com/gin-gonic/gin"
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
	frontPrefix := "/api/v1"
	PublicRouter = Engine.Group(frontPrefix)
	AuthRouter = Engine.Group(frontPrefix)
	AuthRouter.Use(middleware.AuthMiddleware())
}
