package config

// SmsConfig 短信配置
type SmsConfig struct {
	Engine    string            `yaml:"engine" json:"engine"`       // 当前使用的短信服务商
	Whitelist []string          `yaml:"whitelist" json:"whitelist"` // 白名单手机号
	FeigeYun  FeigeYunSmsConfig `yaml:"feigeyun" json:"feigeyun"`
	Aliyun    AliyunSmsConfig   `yaml:"aliyun" json:"aliyun"`
}

// FeigeYunSmsConfig 飞鸽云短信配置
type FeigeYunSmsConfig struct {
	APIKey                   string `yaml:"api-key" json:"api_key"`
	Secret                   string `yaml:"secret" json:"secret"`
	SignID                   string `yaml:"sign-id" json:"sign_id"`
	TemplateID               string `yaml:"template-id" json:"template_id"`
	CreditApprovedTemplateID string `yaml:"credit-approved-template-id" json:"credit_approved_template_id"`
	APIURL                   string `yaml:"api-url" json:"api_url"`
}

// AliyunSmsConfig 阿里云短信配置
type AliyunSmsConfig struct {
	AccessKeyID     string `yaml:"access-key-id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access-key-secret" json:"access_key_secret"`
	SignName        string `yaml:"sign-name" json:"sign_name"`
	TemplateCode    string `yaml:"template-code" json:"template_code"`
}

// SmsCodeConfig 短信验证码配置
type SmsCodeConfig struct {
	CodeLength     int      `yaml:"code-length" json:"code_length"`           // 验证码长度
	CodeExpireTime int      `yaml:"code-expire-time" json:"code_expire_time"` // 验证码过期时间（分钟）
	DailyLimit     int      `yaml:"daily-limit" json:"daily_limit"`           // 每日发送限制
	IntervalLimit  int      `yaml:"interval-limit" json:"interval_limit"`     // 发送间隔限制（分钟）
	RecentLimit    int      `yaml:"recent-limit" json:"recent_limit"`         // 最近N分钟内发送限制次数
	RecentMinutes  int      `yaml:"recent-minutes" json:"recent_minutes"`     // 最近限制时间窗口（分钟）
	Whitelist      []string `yaml:"whitelist" json:"whitelist"`               // 白名单手机号
	WhitelistCode  string   `yaml:"whitelist-code" json:"whitelist_code"`     // 白名单固定验证码
}
