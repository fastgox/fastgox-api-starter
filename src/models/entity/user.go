package entity

import (
	"database/sql/driver"
	"errors"
)

// JSONMap 自定义JSON类型，支持comparable
type JSONMap string

func (j JSONMap) Value() (driver.Value, error) {
	if j == "" {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*j = JSONMap(v)
	case string:
		*j = JSONMap(v)
	default:
		return errors.New("invalid type for JSONMap")
	}
	return nil
}

func (j JSONMap) MarshalJSON() ([]byte, error) {
	if j == "" {
		return []byte("null"), nil
	}
	return []byte(j), nil
}

func (j *JSONMap) UnmarshalJSON(data []byte) error {
	*j = JSONMap(data)
	return nil
}

// User 用户表实体 - 对应 public.user 表
type User struct {
	ID              string  `gorm:"column:id;primaryKey;type:text" json:"id"`
	Name            string  `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email           string  `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Role            string  `gorm:"column:role;type:varchar(255);not null" json:"role"`
	ProfileImageURL string  `gorm:"column:profile_image_url;type:text;not null" json:"profile_image_url"`
	ApiKey          *string `gorm:"column:api_key;type:varchar(255)" json:"api_key,omitempty"`
	CreatedAt       int64   `gorm:"column:created_at;type:bigint;not null" json:"created_at"`
	UpdatedAt       int64   `gorm:"column:updated_at;type:bigint;not null" json:"updated_at"`
	LastActiveAt    int64   `gorm:"column:last_active_at;type:bigint;not null" json:"last_active_at"`
	Settings        JSONMap `gorm:"column:settings;type:json" json:"settings,omitempty"`
	Info            JSONMap `gorm:"column:info;type:json" json:"info,omitempty"`
	OauthSub        *string `gorm:"column:oauth_sub;type:text" json:"oauth_sub,omitempty"`
	InvitationCode  string  `gorm:"column:invitation_code;type:varchar(32);not null;default:''" json:"invitation_code"`
	Phone           *string `gorm:"column:phone;type:varchar(20)" json:"phone,omitempty"`
	OauthData       JSONMap `gorm:"column:oauth_data;type:jsonb" json:"oauth_data,omitempty"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}
