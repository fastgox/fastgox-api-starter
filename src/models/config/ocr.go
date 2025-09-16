package config

// OcrConfig OCR配置
type OcrConfig struct {
	Engine  string           `yaml:"engine" json:"engine"` // 当前使用的OCR服务商
	Aliyun  AliyunOcrConfig  `yaml:"aliyun" json:"aliyun"`
	Qiandun QiandunOcrConfig `yaml:"qiandun" json:"qiandun"`
}

// AliyunOcrConfig 阿里云OCR配置
type AliyunOcrConfig struct {
	AccessKeyID     string `yaml:"access-key-id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access-key-secret" json:"access_key_secret"`
	Endpoint        string `yaml:"endpoint" json:"endpoint"`
	IsProd          bool   `yaml:"is-prod" json:"is_prod"`
}

// QiandunOcrConfig 千盾OCR配置
type QiandunOcrConfig struct {
	AppKey    string `yaml:"app-key" json:"app_key"`
	AppSecret string `yaml:"app-secret" json:"app_secret"`
	BaseURL   string `yaml:"base-url" json:"base_url"`
	IsProd    bool   `yaml:"is-prod" json:"is_prod"`
}
