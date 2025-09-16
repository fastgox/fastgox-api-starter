package config

// AppConfig 应用配置
type AppConfig struct {
	Name    string        `yaml:"name"`
	Version string        `yaml:"version"`
	Env     string        `yaml:"env"`
	Port    int           `yaml:"port"`
	Debug   bool          `yaml:"debug"`
	Swagger SwaggerConfig `yaml:"swagger"`
}

// SwaggerConfig Swagger配置
type SwaggerConfig struct {
	Host     string `yaml:"host"`      // swagger文档显示的host，为空时自动生成
	Scheme   string `yaml:"scheme"`    // http 或 https，为空时根据环境自动判断
	BasePath string `yaml:"base_path"` // API基础路径，默认/api/v1
}
