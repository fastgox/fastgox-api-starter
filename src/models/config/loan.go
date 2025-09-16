package config

// LoanConfig 贷款相关配置
type LoanConfig struct {
	BaseURL string `yaml:"base_url" json:"base_url"` // 贷款成功后跳转的基础URL
}
