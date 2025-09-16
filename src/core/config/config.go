package config

import (
	"fmt"
	"os"

	"github.com/fastgox/fastgox-api-starter/src/models/config"
	"gopkg.in/yaml.v3"
)

// 全局配置实例
var GlobalConfig *config.Config

// InitConfig 初始化配置
func InitConfig() error {
	// 设置环境
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// 构建配置文件路径
	configFile := fmt.Sprintf("config/%s.yaml", env)

	// 检查配置文件是否存在
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在: %s", configFile)
	}

	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析YAML
	var cfg config.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Println(err)
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证配置
	if err := validateConfig(&cfg); err != nil {
		return fmt.Errorf("配置验证失败: %w", err)
	}

	// 设置全局配置
	GlobalConfig = &cfg
	return nil
}

// validateConfig 验证配置
func validateConfig(c *config.Config) error {
	if c.Database.Host == "" {
		return fmt.Errorf("数据库主机地址不能为空")
	}
	if c.Database.User == "" {
		return fmt.Errorf("数据库用户名不能为空")
	}
	if c.Database.DBName == "" {
		return fmt.Errorf("数据库名称不能为空")
	}
	if c.JWT.SecretKey == "" {
		return fmt.Errorf("JWT密钥不能为空")
	}
	if c.App.Port <= 0 || c.App.Port > 65535 {
		return fmt.Errorf("端口号必须在1-65535之间")
	}
	return nil
}
