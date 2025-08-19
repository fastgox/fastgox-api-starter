package database

import (
	"fmt"
	"sync"

	"github.com/fastgox/fastgox-api-starter/src/config"
	"github.com/fastgox/utils/logger"

	"gorm.io/gorm"
)

var (
	globalDB *gorm.DB
	once     sync.Once
)

// Initialize 初始化数据库连接
func Initialize() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		logger.Info("初始化数据库连接...")
		config.InitConfig()
		if config.GlobalConfig == nil {
			err = fmt.Errorf("全局配置未初始化")
			return
		}

		dbConfig := &Config{
			Driver:      config.GlobalConfig.Database.Driver,
			Host:        config.GlobalConfig.Database.Host,
			Port:        config.GlobalConfig.Database.Port,
			User:        config.GlobalConfig.Database.User,
			Password:    config.GlobalConfig.Database.Password,
			DBName:      config.GlobalConfig.Database.DBName,
			SSLMode:     config.GlobalConfig.Database.SSLMode,
			Timezone:    config.GlobalConfig.Database.Timezone,
			MaxOpenConn: config.GlobalConfig.Database.MaxOpenConn,
			MaxIdleConn: config.GlobalConfig.Database.MaxIdleConn,
			LogLevel:    config.GlobalConfig.Database.LogLevel,
		}

		globalDB, err = NewConnection(dbConfig)
		if err != nil {
			err = fmt.Errorf("数据库连接失败: %w", err)
			return
		}

		logger.Info("数据库初始化完成")
	})

	return globalDB, err
}

// GetDB 获取全局数据库实例
func GetDB() *gorm.DB {
	if globalDB == nil {
		panic("数据库未初始化，请先调用 Initialize()")
	}
	return globalDB
}

// AutoMigrate 自动迁移表结构
func AutoMigrate(models ...interface{}) error {
	if globalDB == nil {
		return fmt.Errorf("数据库未初始化")
	}

	return globalDB.AutoMigrate(models...)
}

// TestConnection 测试数据库连接
func TestConnection() error {
	if globalDB == nil {
		return fmt.Errorf("数据库未初始化")
	}

	sqlDB, err := globalDB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %w", err)
	}

	return sqlDB.Ping()
}

// GetStats 获取数据库连接统计信息
func GetStats() (map[string]interface{}, error) {
	if globalDB == nil {
		return nil, fmt.Errorf("数据库未初始化")
	}

	sqlDB, err := globalDB.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %w", err)
	}

	stats := sqlDB.Stats()
	return map[string]interface{}{
		"open_connections":     stats.OpenConnections,
		"in_use":               stats.InUse,
		"idle":                 stats.Idle,
		"wait_count":           stats.WaitCount,
		"wait_duration":        stats.WaitDuration.String(),
		"max_idle_closed":      stats.MaxIdleClosed,
		"max_idle_time_closed": stats.MaxIdleTimeClosed,
		"max_lifetime_closed":  stats.MaxLifetimeClosed,
	}, nil
}

// Close 关闭数据库连接
func Close() error {
	if globalDB == nil {
		return nil
	}

	sqlDB, err := globalDB.DB()
	if err != nil {
		return err
	}

	logger.Info("关闭数据库连接...")
	return sqlDB.Close()
}
