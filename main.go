// Package main provides the entry point for the fastgox-api-starter API server
//
//	@title						fastgox-api-starter API
//	@version					1.0
//	@description				fastgox-api-starter 贷款申请服务API
//	@BasePath					/api/v1
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/fastgox/fastgox-api-starter/docs"
	"github.com/fastgox/fastgox-api-starter/src"
	"github.com/fastgox/utils/logger"
)

func main() {
	logger.Info("启动 fastgox-api-starter 服务...")

	// 创建服务器实例
	srv, err := src.NewServer()
	if err != nil {
		logger.Error("创建服务器失败: %v", err)
		os.Exit(1)
	}

	// 启动服务
	if err := srv.Start(); err != nil {
		logger.Error("启动服务器失败: %v", err)
		os.Exit(1)
	}

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("🛑 正在关闭服务...")

	// 优雅关闭服务
	if err := srv.Stop(); err != nil {
		logger.Error("服务器关闭失败: %v", err)
	}

	logger.Info("服务器已安全关闭")
}
