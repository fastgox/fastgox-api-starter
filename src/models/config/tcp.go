package config

// TCPConfig TCP服务配置
type TCPConfig struct {
	Enabled          bool `yaml:"enabled"`
	Port             int  `yaml:"port"`
	MaxConnections   int  `yaml:"max_connections"`
	ReadTimeout      int  `yaml:"read_timeout"`      // 读超时(秒)
	WriteTimeout     int  `yaml:"write_timeout"`     // 写超时(秒)
	MaxMessageSize   int  `yaml:"max_message_size"`  // 最大消息体(字节)
	HeartbeatTimeout int  `yaml:"heartbeat_timeout"` // 心跳超时(秒)，超过此时间没收到消息则断开，0=不启用
}
