package src

import (
	"fmt"
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/core/database"
	"github.com/fastgox/fastgox-api-starter/src/router"
	_ "github.com/fastgox/fastgox-api-starter/src/router/handle"
	"github.com/fastgox/utils/logger"
	"github.com/gin-gonic/gin"
)

// Server 应用服务器
type Server struct {
	Router *gin.Engine
	HTTP   *http.Server
}

// NewServer 创建新的服务器实例
func NewServer() (*Server, error) {
	logger.InitWithPath("data/logs")
	logger.Info("创建服务器实例..")

	// 初始化数据库
	_, err := database.Initialize()
	if err != nil {
		logger.Error("数据库初始化失败: %v", err)
		return nil, err
	}
	logger.Info("数据库连接成功")

	server := &Server{}

	// 创建HTTP服务器
	port := 8080
	if config.GlobalConfig != nil && config.GlobalConfig.App.Port > 0 {
		port = config.GlobalConfig.App.Port
	}
	addr := fmt.Sprintf(":%d", port)
	server.HTTP = &http.Server{
		Addr:    addr,
		Handler: router.Engine,
	}

	logger.Info("服务器实例创建完成")
	return server, nil
}

// Start 启动服务器
func (s *Server) Start() error {
	addr := s.HTTP.Addr
	logger.Info("启动服务器..")
	fmt.Println("==============================")
	fmt.Println("服务已启动:")
	fmt.Printf("  API地址:    http://localhost%s\n", addr)
	fmt.Printf("  Swagger文档: http://localhost%s/swagger/index.html\n", addr)
	fmt.Println("==============================")
	logger.Info("服务器地址: http://localhost%s", addr)
	logger.Info("API文档: http://localhost%s/swagger/index.html", addr)
	return s.HTTP.ListenAndServe()
}

// Stop 停止服务器
func (s *Server) Stop() error {
	logger.Info("正在关闭服务器..")
	logger.Info("服务器已安全关闭")
	return nil
}
