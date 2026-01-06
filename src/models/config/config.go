package config

// Config 应用配置结构体
type Config struct {
	App      AppConfig    `yaml:"app"`
	Database DatabaseConf `yaml:"database"`
	JWT      JWTConfig    `yaml:"jwt"`
}
