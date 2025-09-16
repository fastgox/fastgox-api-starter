package config

// DatabaseConf 数据库配置
type DatabaseConf struct {
	Driver      string `yaml:"driver"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DBName      string `yaml:"dbname"`
	SSLMode     string `yaml:"sslmode"`
	Timezone    string `yaml:"timezone"`
	MaxOpenConn int    `yaml:"max_open_conns"`
	MaxIdleConn int    `yaml:"max_idle_conns"`
	LogLevel    string `yaml:"log_level"`
}
