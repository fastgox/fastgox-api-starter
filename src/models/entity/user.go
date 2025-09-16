package entity

import "time"

// User 用户表实体 - 对应 t_user 表
type User struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Phone       string     `gorm:"column:phone;type:varchar(80)" json:"phone"`
	ChannelCode string     `gorm:"column:channel_code;type:varchar(80)" json:"channel_code"` // 渠道代码，替代channel_id
	Status      int8       `gorm:"column:status;type:tinyint(1);default:1" json:"status"`
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime(3)" json:"create_time"`
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime(3)" json:"update_time"`
	Platform    string     `gorm:"column:platform;type:varchar(80)" json:"platform"`
	IsAuth      int8       `gorm:"column:is_auth;type:tinyint;default:0" json:"is_auth"`
}

// TableName 设置表名
func (User) TableName() string {
	return "t_user"
}

// IsAuthenticated 判断用户是否已实名认证
func (u *User) IsAuthenticated() bool {
	return u.IsAuth == 1
}

// IsActive 判断用户是否激活状态
func (u *User) IsActive() bool {
	return u.Status == 1
}
