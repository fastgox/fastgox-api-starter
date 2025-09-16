package config

// AuthConfig 身份认证配置
type AuthConfig struct {
	Engine  string            `yaml:"engine" json:"engine"` // 当前使用的身份认证服务商
	Qiandun QiandunAuthConfig `yaml:"qiandun" json:"qiandun"`
}

// QiandunAuthConfig 钱盾认证配置
type QiandunAuthConfig struct {
	AppKey    string `yaml:"app-key" json:"app_key"`
	AppSecret string `yaml:"app-secret" json:"app_secret"`
	IsProd    bool   `yaml:"is-prod" json:"is_prod"`
}
