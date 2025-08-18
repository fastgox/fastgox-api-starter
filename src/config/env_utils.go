package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fastgox/utils/logger"
)

// LoadEnvString 加载字符串环境变量
func LoadEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		logger.Debug(fmt.Sprintf("环境变量 %s 已设置", key))
		return value
	}
	logger.Debug(fmt.Sprintf("环境变量 %s 未设置，使用默认值 %s", key, defaultValue))
	return defaultValue
}

// LoadEnvInt 加载整数环境变量
func LoadEnvInt(key string, defaultValue int) int {
	if valueStr := os.Getenv(key); valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			logger.Debug(fmt.Sprintf("环境变量 %s 已设置为: %d", key, value))
			return value
		} else {
			logger.Warn(fmt.Sprintf("环境变量 %s 格式错误，使用默认值 %d", key, defaultValue))
		}
	}
	logger.Debug(fmt.Sprintf("环境变量 %s 未设置，使用默认值 %d", key, defaultValue))
	return defaultValue
}

// LoadEnvBool 加载布尔环境变量
func LoadEnvBool(key string, defaultValue bool) bool {
	if valueStr := os.Getenv(key); valueStr != "" {
		valueStr = strings.ToLower(valueStr)
		switch valueStr {
		case "true", "1", "yes", "on":
			logger.Debug(fmt.Sprintf("环境变量 %s 已设置为: true", key))
			return true
		case "false", "0", "no", "off":
			logger.Debug(fmt.Sprintf("环境变量 %s 已设置为: false", key))
			return false
		default:
			logger.Warn(fmt.Sprintf("环境变量 %s 格式错误，使用默认值 %t", key, defaultValue))
		}
	}
	logger.Debug(fmt.Sprintf("环境变量 %s 未设置，使用默认值 %t", key, defaultValue))
	return defaultValue
}
