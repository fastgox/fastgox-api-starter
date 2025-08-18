package config

// Config 应用配置结构体
type Config struct {
	App      AppConfig    `yaml:"app"`
	Database DatabaseConf `yaml:"database"`
	JWT      JWTConfig    `yaml:"jwt"`
}

type JWTConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type DatabaseConf struct {
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

type AppConfig struct {
	Name    string        `yaml:"name"`
	Version string        `yaml:"version"`
	Env     string        `yaml:"env"`
	Port    int           `yaml:"port"`
	Debug   bool          `yaml:"debug"`
	Swagger SwaggerConfig `yaml:"swagger"`
}

type SwaggerConfig struct {
	Host     string `yaml:"host"`      // swagger文档显示的host，为空时自动生成
	Scheme   string `yaml:"scheme"`    // http 或 https，为空时根据环境自动判断
	BasePath string `yaml:"base_path"` // API基础路径，默认/api/v1
}
