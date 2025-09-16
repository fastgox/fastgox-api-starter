package config

// Config 应用配置结构体
type Config struct {
	App         AppConfig         `yaml:"app"`
	Database    DatabaseConf      `yaml:"database"`
	JWT         JWTConfig         `yaml:"jwt"`
	SMS         SmsConfig         `yaml:"sms"`
	OCR         OcrConfig         `yaml:"ocr"`
	Auth        AuthConfig        `yaml:"auth"`
	SmsCode     SmsCodeConfig     `yaml:"sms-code"`
	AppMarket   AppMarketConfig   `yaml:"appmarket"`
	Geolocation GeolocationConfig `yaml:"geolocation"`
	File        FileConfig        `yaml:"file"`
	Loan        LoanConfig        `yaml:"loan"`
}
