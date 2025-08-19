package main

import (
	"fmt"
	"os"

	"github.com/fastgox/fastgox-api-starter/src/config"
	"github.com/fastgox/fastgox-api-starter/src/core/database"
	"github.com/fastgox/utils/logger"
)

func main() {
	logger.InitWithPath("data/logs")
	logger.Info("测试MySQL数据库连接...")

	// 设置环境变量使用MySQL配置
	os.Setenv("APP_ENV", "mysql")

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		logger.Error("配置初始化失败: ", err)
		fmt.Printf("❌ 配置初始化失败: %v\n", err)
		return
	}

	// 创建数据库配置
	dbConfig := &database.Config{
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

	fmt.Printf("数据库驱动: %s\n", dbConfig.Driver)
	fmt.Printf("连接地址: %s:%d\n", dbConfig.Host, dbConfig.Port)
	fmt.Printf("数据库名: %s\n", dbConfig.DBName)
	fmt.Printf("DSN: %s\n", dbConfig.DSN())

	// 测试连接
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		logger.Error("MySQL数据库连接失败: ", err)
		fmt.Printf("❌ MySQL数据库连接失败: %v\n", err)
		return
	}

	// 测试ping
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("获取数据库实例失败: ", err)
		fmt.Printf("❌ 获取数据库实例失败: %v\n", err)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Error("数据库ping失败: ", err)
		fmt.Printf("❌ 数据库ping失败: %v\n", err)
		return
	}

	logger.Info("✅ MySQL数据库连接成功！")
	fmt.Println("✅ MySQL数据库连接测试通过")

	// 关闭连接
	sqlDB.Close()
}
