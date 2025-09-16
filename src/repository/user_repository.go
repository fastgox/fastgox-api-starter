package repository

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// UserRepository 用户仓储
type UserRepository struct {
	BaseRepository[entity.User]
}

var UserRepo = &UserRepository{BaseRepository: *GetRepository[entity.User]()}

// GetByPhone 根据手机号获取用户
func (r *UserRepository) GetByPhone(phone string) (*entity.User, error) {
	return r.First("phone = ?", phone)
}

// FindOrCreateUserByPhone 根据手机号查找或创建用户
func (r *UserRepository) FindOrCreateUserByPhone(phone string) (*entity.User, bool, error) {
	// 先尝试查找现有用户
	user, err := r.GetByPhone(phone)
	if err == nil && user != nil {
		return user, false, nil // 找到用户，不是新用户
	}

	// 用户不存在，创建新用户
	now := time.Now()
	newUser := &entity.User{
		Phone:       phone,
		ChannelCode: "sms_login",
		Status:      1,
		CreateTime:  &now,
		UpdateTime:  &now,
		Platform:    "mobile",
		IsAuth:      0,
	}

	if err := r.Create(newUser); err != nil {
		return nil, false, err
	}

	return newUser, true, nil // 创建成功，是新用户
}

// UpdateLoginTime 更新用户登录时间
func (r *UserRepository) UpdateLoginTime(userID int64) error {
	now := time.Now()
	return r.DB.Model(&entity.User{}).
		Where("id = ?", userID).
		Update("update_time", now).Error
}

// UpdateAuthStatus 更新用户认证状态
func (r *UserRepository) UpdateAuthStatus(userID int64, isAuth int8) error {
	now := time.Now()
	return r.DB.Model(&entity.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"is_auth":     isAuth,
			"update_time": now,
		}).Error
}
