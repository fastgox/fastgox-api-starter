package entity

// User 用户表实体
type User struct {
	BaseModel
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Email string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Role  string `gorm:"type:varchar(50);not null;default:'user'" json:"role"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}

// IsAdmin 判断是否为管理员
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
