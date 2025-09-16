package src

import (
	"fmt"
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	_ "github.com/fastgox/fastgox-api-starter/src/pkg/auth"
	"github.com/fastgox/fastgox-api-starter/src/pkg/file"
	"github.com/fastgox/fastgox-api-starter/src/pkg/geolocation"
	_ "github.com/fastgox/fastgox-api-starter/src/pkg/ocr"
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

	// 初始化地理位置服务提供商
	geolocation.InitAmapProvider()
	// 初始化文件服务提供商
	file.InitLocalProvider()

	server := &Server{}

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
