package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// getLogLevel 根据字符串获取GORM日志级别
func getLogLevel(level string) gormlogger.LogLevel {
	switch level {
	case "debug":
		return gormlogger.Info
	case "info":
		return gormlogger.Warn
	case "warn":
		return gormlogger.Error
	case "error":
		return gormlogger.Error
	default:
		return gormlogger.Silent
	}
}

// NewConnection 创建数据库连接
func NewConnection(config *Config) (*gorm.DB, error) {
	if config == nil {
		return nil, fmt.Errorf("配置不能为空")
	}

	// 验证配置
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("配置验证失败: %w", err)
	}

	// 根据数据库类型选择驱动
	var dialector gorm.Dialector
	switch config.Driver {
	case "mysql":
		dialector = mysql.Open(config.DSN())
	case "postgres", "postgresql":
		dialector = postgres.Open(config.DSN())
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", config.Driver)
	}

	// 连接数据库
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormlogger.Default.LogMode(getLogLevel(config.LogLevel)),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接测试失败: %v", err)
	}

	return db, nil
}
