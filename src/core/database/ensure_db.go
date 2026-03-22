package database

import (
	"fmt"
	"net/url"

	"github.com/fastgox/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// EnsureDatabase 确保目标数据库存在，不存在则自动创建
func EnsureDatabase(config *Config) error {
	dbName := config.DBName
	if dbName == "" {
		return fmt.Errorf("数据库名称不能为空")
	}

	logger.Info("[Migration] 检查数据库 '%s' 是否存在...", dbName)

	// 1. 连接到系统数据库
	sysDB, err := connectSystemDB(config)
	if err != nil {
		return fmt.Errorf("连接系统数据库失败: %w", err)
	}
	defer func() {
		if sqlDB, err := sysDB.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// 2. 检查目标数据库是否存在
	exists, err := databaseExists(sysDB, config.Driver, dbName)
	if err != nil {
		return fmt.Errorf("检查数据库失败: %w", err)
	}

	if exists {
		logger.Info("[Migration] 数据库 '%s' 已存在", dbName)
		return nil
	}

	// 3. 创建数据库
	logger.Info("[Migration] 数据库 '%s' 不存在，正在创建...", dbName)
	if err := createDatabase(sysDB, config.Driver, dbName); err != nil {
		return fmt.Errorf("创建数据库失败: %w", err)
	}

	logger.Info("[Migration] ✅ 数据库 '%s' 创建成功", dbName)
	return nil
}

// connectSystemDB 连接到系统数据库（不指定目标数据库）
func connectSystemDB(config *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch config.Driver {
	case "mysql":
		// MySQL: 连接时不指定数据库名
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=%s&timeout=30s&allowNativePasswords=true&tls=false",
			config.User, config.Password, config.Host, config.Port,
			url.QueryEscape(config.Timezone),
		)
		dialector = mysql.Open(dsn)

	case "postgres", "postgresql":
		// PostgreSQL: 连接到默认的 "postgres" 数据库
		escapedPassword := url.QueryEscape(config.Password)
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=postgres port=%d sslmode=%s TimeZone=%s",
			config.Host, config.User, escapedPassword, config.Port, config.SSLMode, config.Timezone,
		)
		dialector = postgres.Open(dsn)

	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", config.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// databaseExists 检查数据库是否存在
func databaseExists(db *gorm.DB, driver, dbName string) (bool, error) {
	var count int64

	switch driver {
	case "mysql":
		err := db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count).Error
		return count > 0, err

	case "postgres", "postgresql":
		err := db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count).Error
		return count > 0, err

	default:
		return false, fmt.Errorf("不支持的数据库驱动: %s", driver)
	}
}

// createDatabase 创建数据库
func createDatabase(db *gorm.DB, driver, dbName string) error {
	switch driver {
	case "mysql":
		sql := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)
		return db.Exec(sql).Error

	case "postgres", "postgresql":
		sql := fmt.Sprintf(`CREATE DATABASE "%s" ENCODING 'UTF8'`, dbName)
		return db.Exec(sql).Error

	default:
		return fmt.Errorf("不支持的数据库驱动: %s", driver)
	}
}
