package config

// AppConfig TCP游戏服务器应用配置
type AppConfig struct {
	Name    string `yaml:"name"`    // 应用名称
	Version string `yaml:"version"` // 应用版本
	Env     string `yaml:"env"`     // 环境：dev, test, prod
	Host    string `yaml:"host"`    // 服务器监听地址
	Port    int    `yaml:"port"`    // 服务器监听端口
	Debug   bool   `yaml:"debug"`   // 是否开启调试模式
}
