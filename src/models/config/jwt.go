package config

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey string `yaml:"secret_key"`
}
