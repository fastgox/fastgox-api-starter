package repository

import (
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// UserRepository 用户仓储
type UserRepository struct {
	BaseRepository[entity.User]
}

var UserRepo = &UserRepository{BaseRepository: *GetRepository[entity.User]()}

// GetByEmail 根据邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	return r.First("email = ?", email)
}
