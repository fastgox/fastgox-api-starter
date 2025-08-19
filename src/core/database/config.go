package database

// Config 数据库配置结构
type Config struct {
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	User        string `json:"user" yaml:"user"`
	Password    string `json:"password" yaml:"password"`
	DBName      string `json:"dbname" yaml:"dbname"`
	SSLMode     string `json:"sslmode" yaml:"sslmode"`
	Timezone    string `json:"timezone" yaml:"timezone"`
	MaxOpenConn int    `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConn int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	LogLevel    string `json:"log_level" yaml:"log_level"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Host:        "localhost",
		Port:        5432,
		User:        "postgres",
		Password:    "",
		DBName:      "hkchat",
		SSLMode:     "disable",
		Timezone:    "Asia/Shanghai",
		MaxOpenConn: 100,
		MaxIdleConn: 10,
		LogLevel:    "warn",
	}
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Host == "" {
		return ErrInvalidHost
	}
	if c.Port <= 0 || c.Port > 65535 {
		return ErrInvalidPort
	}
	if c.User == "" {
		return ErrInvalidUser
	}
	if c.DBName == "" {
		return ErrInvalidDBName
	}
	if c.MaxOpenConn <= 0 {
		c.MaxOpenConn = 100
	}
	if c.MaxIdleConn <= 0 {
		c.MaxIdleConn = 10
	}
	return nil
}

// DSN 获取数据库连接字符串
func (c *Config) DSN() string {
	return formatDSN(c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode, c.Timezone)
}
