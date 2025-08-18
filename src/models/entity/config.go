package entity

// Config config 表实体
type Config struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Version   int64  `gorm:"not null" json:"version"`
	CreatedAt int64  `gorm:"not null" json:"created_at"`
	UpdatedAt *int64 `gorm:"-:all" json:"updated_at"` // 禁用GORM的自动管理
	Key       string `gorm:"not null" json:"key"`
	Val       string `gorm:"not null" json:"val"`
	Module    string `gorm:"not null;column:module" json:"module"`
	//key的字段类型
	FieldType string `gorm:"not null" json:"field_type"`
}

// TableName 设置表名
func (Config) TableName() string {
	return "config_go"
}

// 实现 ConfigRecord 接口
func (c Config) GetKey() string       { return c.Key }
func (c Config) GetVal() string       { return c.Val }
func (c Config) GetFieldType() string { return c.FieldType }
