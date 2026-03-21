package config

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
