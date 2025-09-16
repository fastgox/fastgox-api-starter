package config

// GeolocationConfig 地理位置服务配置
type GeolocationConfig struct {
	DefaultProvider string                       `yaml:"default_provider"` // 默认提供商
	Providers       map[string]GeolocationAPIKey `yaml:"providers"`        // 提供商配置
}

// GeolocationAPIKey 地理位置服务API密钥配置
type GeolocationAPIKey struct {
	APIKey string `yaml:"api_key"` // API密钥
}
