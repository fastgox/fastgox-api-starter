package src

import (
	"fmt"

	coreConfig "github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/router"
	_ "github.com/fastgox/fastgox-api-starter/src/router/handle" // 自动导入所有handle进行注册
	"github.com/fastgox/utils/logger"
	"github.com/lonng/nano"
)

// Server TCP游戏服务器（基于nano框架）
type Server struct {
	addr    string
	enabled bool
}

// NewServer 创建新的TCP游戏服务器实例
func NewServer() (*Server, error) {
	// 初始化日志系统
	err := logger.InitWithPath("data/logs")
	if err != nil {
		return nil, fmt.Errorf("初始化日志系统失败: %w", err)
	}

	logger.Info("创建TCP游戏服务器实例...")

	// 初始化配置
	if err := coreConfig.InitConfig(); err != nil {
		return nil, fmt.Errorf("初始化配置失败: %w", err)
	}

	addr := fmt.Sprintf("%s:%d", coreConfig.GlobalConfig.App.Host, coreConfig.GlobalConfig.App.Port)

	server := &Server{
		addr:    addr,
		enabled: true, // 默认启用TCP游戏服务器
	}

	logger.Info("TCP游戏服务器实例创建完成")
	return server, nil
}

// Start 启动TCP游戏服务器
func (s *Server) Start() error {
	if !s.enabled {
		logger.Info("游戏服务器未启用")
		return nil
	}

	logger.Info("启动TCP游戏服务器...")
	logger.Info("🎮 游戏服务器地址: %s", s.addr)
	logger.Info("🚀 基于nano框架的TCP游戏服务器")

	// 获取已自动注册的TCP组件
	components := router.GetComponents()
	logger.Info("TCP组件已通过init函数自动注册完成")

	logger.Info("nano服务器启动中...")
	logger.Info("✅ nano TCP游戏框架集成完成！")
	logger.Info("🎮 WebSocket连接地址: ws://%s/ws", s.addr)

	// 启动nano服务器（这会阻塞）
	nano.Listen(s.addr,
		nano.WithComponents(components),
		nano.WithIsWebsocket(true),
	)

	return nil
}

// Stop 停止TCP游戏服务器
func (s *Server) Stop() error {
	if !s.enabled {
		return nil
	}

	logger.Info("正在关闭TCP游戏服务器...")
	nano.Shutdown()
	logger.Info("TCP游戏服务器已安全关闭")

	// 关闭所有logger
	logger.CloseAll()
	return nil
}
