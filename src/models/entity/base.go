package entity

// BaseModel 基础模型，包含ID、创建时间和更新时间
type BaseModel struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt int64  `gorm:"type:bigint;column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;column:updated_at" json:"updated_at"`
}
