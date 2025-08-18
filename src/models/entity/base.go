package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel 基础模型，包含ID、创建时间和更新时间
type BaseModel struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt int64  `gorm:"type:bigint;column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;column:updated_at" json:"updated_at"`
}

// BeforeCreate GORM 钩子：创建前自动生成UUID和设置时间戳
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Unix()

	// 自动生成UUID（如果ID为空）
	if b.ID == "" {
		b.ID = uuid.New().String()
	}

	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate GORM 钩子：更新前自动设置时间戳
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now().Unix()
	return nil
}

// BaseModelWithoutUpdate 基础模型（仅创建时间），用于不需要更新时间的表
type BaseModelWithoutUpdate struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt int64  `gorm:"type:bigint;column:created_at" json:"created_at"`
}

// BeforeCreate GORM 钩子：创建前自动生成UUID和设置时间戳
func (b *BaseModelWithoutUpdate) BeforeCreate(tx *gorm.DB) error {
	// 自动生成UUID（如果ID为空）
	if b.ID == "" {
		b.ID = uuid.New().String()
	}

	b.CreatedAt = time.Now().Unix()
	return nil
}

// BaseModelNano 基础模型（纳秒时间戳），用于需要更高精度时间的表（如消息）
type BaseModelNano struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt int64  `gorm:"type:bigint;column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;column:updated_at" json:"updated_at"`
}

// BeforeCreate GORM 钩子：创建前自动生成UUID和设置纳秒时间戳
func (b *BaseModelNano) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UnixNano()

	// 自动生成UUID（如果ID为空）
	if b.ID == "" {
		b.ID = uuid.New().String()
	}

	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate GORM 钩子：更新前自动设置纳秒时间戳
func (b *BaseModelNano) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now().UnixNano()
	return nil
}

// BaseModelNanoWithoutUpdate 基础模型（纳秒时间戳，仅创建时间）
type BaseModelNanoWithoutUpdate struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt int64  `gorm:"type:bigint;column:created_at" json:"created_at"`
}

// BeforeCreate GORM 钩子：创建前自动生成UUID和设置纳秒时间戳
func (b *BaseModelNanoWithoutUpdate) BeforeCreate(tx *gorm.DB) error {
	// 自动生成UUID（如果ID为空）
	if b.ID == "" {
		b.ID = uuid.New().String()
	}

	b.CreatedAt = time.Now().UnixNano()
	return nil
}
