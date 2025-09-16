package config

// AppMarketConfig 应用市场配置
type AppMarketConfig struct {
	Engine             string                `yaml:"engine" json:"engine"`                           // 当前使用的应用市场服务商
	SupportedPlatforms []string              `yaml:"supported_platforms" json:"supported_platforms"` // 支持的平台列表
	Xiaomi             XiaomiAppMarketConfig `yaml:"xiaomi" json:"xiaomi"`
	Vivo               VivoAppMarketConfig   `yaml:"vivo" json:"vivo"`
}

// XiaomiAppMarketConfig 小米应用市场配置
type XiaomiAppMarketConfig struct {
	APIKey    string `yaml:"api-key" json:"api_key"`
	SecretKey string `yaml:"secret-key" json:"secret_key"`
	BaseURL   string `yaml:"base-url" json:"base_url"`
	IsProd    bool   `yaml:"is-prod" json:"is_prod"`
}

// VivoAppMarketConfig vivo应用市场配置
type VivoAppMarketConfig struct {
	APIKey    string `yaml:"api-key" json:"api_key"`
	SecretKey string `yaml:"secret-key" json:"secret_key"`
	BaseURL   string `yaml:"base-url" json:"base_url"`
	IsProd    bool   `yaml:"is-prod" json:"is_prod"`
}
