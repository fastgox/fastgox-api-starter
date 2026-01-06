package entity

// Auth 认证表实体 - 对应 public.auth 表
type Auth struct {
	ID       string  `gorm:"column:id;primaryKey;type:text" json:"id"`
	Email    string  `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Password string  `gorm:"column:password;type:text;not null" json:"-"`
	Active   bool    `gorm:"column:active;type:bool;not null" json:"active"`
	Phone    *string `gorm:"column:phone;type:varchar(20)" json:"phone,omitempty"`
}

// TableName 设置表名
func (Auth) TableName() string {
	return "auth"
}
