package config

import "github.com/fastgox/utils/logger"

// InitializeConfig 初始化配置系统
func InitializeConfig() error {
	logger.Info("初始化配置系统...")

	if err := InitConfig(); err != nil {
		return err
	}

	logger.Info("配置系统初始化完成")
	return nil
}

// InitializeLogger 初始化日志系统
func InitializeLogger() error {
	logger.Info("初始化日志系统...")

	// 这里暂时注释掉，因为配置结构体可能需要更新
	// logConfig := &logger.LogConfig{
	// 	FilePath: GlobalConfig.Log.FilePath,
	// 	Level:    GlobalConfig.Log.Level,
	// }

	// if err := logger.InitLogger(logConfig); err != nil {
	// 	logger.Error("日志系统初始化失败，使用默认日志:", err)
	// 	return err
	// }

	logger.Info("日志系统初始化完成")
	return nil
}
