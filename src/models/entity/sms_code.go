package entity

import (
	"time"

	"gorm.io/gorm"
)

// SmsCode 短信验证码实体
// 对应数据库表 t_sms_codes
// @Description 短信验证码信息
type SmsCode struct {
	ID         uint           `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt  *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
	Phone      string         `gorm:"column:phone" json:"phone"`               // 手机号
	Code       string         `gorm:"column:code" json:"code"`                 // 验证码
	ExpireTime *time.Time     `gorm:"column:expire_time" json:"expire_time"`   // 过期时间
	IsUsed     bool           `gorm:"column:is_used;default:0" json:"is_used"` // 是否已使用
}

// TableName 指定表名
func (SmsCode) TableName() string {
	return "t_sms_codes"
}

// IsExpired 检查验证码是否过期
func (s *SmsCode) IsExpired() bool {
	if s.ExpireTime == nil {
		return true
	}
	return time.Now().After(*s.ExpireTime)
}

// IsValid 检查验证码是否有效（未使用且未过期）
func (s *SmsCode) IsValid() bool {
	return !s.IsUsed && !s.IsExpired()
}
