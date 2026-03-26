package src

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg/file"
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

// Start 非阻塞启动服务器，在后台监听请求
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

	// 非阻塞启动，让 main 可以继续监听信号
	go func() {
		if err := s.HTTP.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("HTTP服务器异常退出: %v", err)
		}
	}()

	return nil
}

// Stop 优雅关闭服务器，等待正在处理的请求完成
func (s *Server) Stop() error {
	timeout := config.GlobalConfig.App.ShutdownTimeout
	if timeout <= 0 {
		timeout = 30 // 默认30秒
	}

	logger.Info("正在优雅关闭服务器，最长等待 %d 秒...", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Shutdown 会停止接收新请求，并等待已有请求处理完毕
	if err := s.HTTP.Shutdown(ctx); err != nil {
		logger.Error("服务器优雅关闭超时，强制退出: %v", err)
		return err
	}

	logger.Info("服务器已安全关闭")
	return nil
}
