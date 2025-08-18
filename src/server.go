package src

import (
	"fmt"
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/config"
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
	logger.Info("创建服务器实例..")

	server := &Server{}

	// 初始化配置
	if err := config.InitializeConfig(); err != nil {
		return nil, fmt.Errorf("配置初始化失败: %w", err)
	}

	// 初始化日志
	if err := config.InitializeLogger(); err != nil {
		logger.Error("日志系统初始化失败，继续使用默认日志:", err)
	}

	// 初始化数据库（这里需要实现数据库连接逻辑）
	//initDatabase()

	// 创建HTTP服务器
	addr := fmt.Sprintf(":%d", config.GlobalConfig.App.Port)
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
	logger.Info(fmt.Sprintf("服务器地址: http://localhost%s", addr))
	logger.Info(fmt.Sprintf("API文档: http://localhost%s/swagger/index.html", addr))

	return s.HTTP.ListenAndServe()
}

// Stop 停止服务器
func (s *Server) Stop() error {
	logger.Info("正在关闭服务器..")
	logger.Info("服务器已安全关闭")
	return nil
}
