package repository

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// UserRepository 用户仓储
type UserRepository struct {
	*BaseRepository[entity.User]
}

var UserRepo = &UserRepository{}

func (r *UserRepository) getRepo() *BaseRepository[entity.User] {
	if r.BaseRepository == nil {
		r.BaseRepository = GetRepository[entity.User]()
	}
	return r.BaseRepository
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id string) (*entity.User, error) {
	return r.getRepo().First("id = ?", id)
}

// GetByEmail 根据邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	return r.getRepo().First("email = ?", email)
}

// GetByPhone 根据手机号获取用户
func (r *UserRepository) GetByPhone(phone string) (*entity.User, error) {
	return r.getRepo().First("phone = ?", phone)
}

// UpdateLastActiveAt 更新最后活跃时间
func (r *UserRepository) UpdateLastActiveAt(id string) error {
	now := time.Now().Unix()
	return r.getRepo().DB.Model(&entity.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"last_active_at": now,
			"updated_at":     now,
		}).Error
}
