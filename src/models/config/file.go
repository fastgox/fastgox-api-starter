package config

// FileConfig 文件配置
type FileConfig struct {
	MaxSize           int64    `yaml:"max_size" json:"max_size"`                     // 最大文件大小（字节）
	MaxCount          int      `yaml:"max_count" json:"max_count"`                   // 单次最大上传文件数量
	AllowedExtensions []string `yaml:"allowed_extensions" json:"allowed_extensions"` // 允许的文件扩展名
	UploadPath        string   `yaml:"upload_path" json:"upload_path"`               // 上传文件存储路径
	URLPrefix         string   `yaml:"url_prefix" json:"url_prefix"`                 // 文件访问URL前缀
}
